import React, {useRef, useState, useEffect} from 'react';
import ProTable, {ProColumns, ActionType} from '@ant-design/pro-table';
import * as dayjs from "dayjs";
import {Button, Divider, message, Popconfirm, Space, Tag} from "antd";
import {DataNode} from "rc-tree-select/lib/interface";
import {deleteRecipe, recipe, saveRecipe, recipes, recipeCategories} from "@/services/recipe";
import RecipeForm from "@/pages/recipe/List/components/RecipeForm";
import {RecipeWeight} from "@/utils/constant";

const RecipeList: React.FC = () => {
  const [formVisible, handleFormVisible]            = useState<boolean>(false)
  const [row, setRow]                               = useState<API.RecipeItem>()
  const [categoriesTreeData, setCategoriesTreeData] = useState<DataNode[]>()
  const actionRef                                   = useRef<ActionType>()
  const weights = [
    {
      value: RecipeWeight.Hot,
      label: "热门",
      color: "red",
    },
    {
      value: RecipeWeight.Favorite,
      label: "最爱",
      color: "magenta",
    }]
  const getCategoriesTreeData = async () => {
    const res = await recipeCategories(0, 0, {all: 1, parent_id: "", not_id: ""});
    if (!res.success) {
      message.error('获取失败');
      return [];
    }

    return res.data.map((r) => {
      return {
        id        : r.id,
        pId       : r.parent_id,
        value     : r.id,
        title     : r.name,
        isLeaf    : r.parent_id !== 0,
        selectable: r.parent_id !== 0
      }
    })
  };

  useEffect(() => {
    (async () => {
      setCategoriesTreeData(await getCategoriesTreeData());
    })()

  }, []);

  const columns: ProColumns<API.RecipeItem>[] = [
    {
      title    : 'ID',
      dataIndex: 'id',
      width    : '10%',
    },
    {
      title    : '名称',
      dataIndex: 'title',
      width    : '20%',
    },
    {
      title    : '分类',
      dataIndex: 'category_id',
      render: (_, record) => categoriesTreeData?.find(e => e.id === record.category_id)?.title

    },
    {
      title    : '排序',
      dataIndex: 'sort',
      search   : false,
    },
    {
      title    : '收藏数',
      search   : false,
      dataIndex: 'favorites',
    },
    {
      title    : '权重',
      dataIndex: 'weight',
      search   : false,
      render: (_, record) => (
        <Space>
          {weights.map((w) => {
            console.log((record.weight || 0 ) & w.value, w.value);


              if (((record.weight || 0 ) & w.value) === w.value ) {
                return (  <Tag  color={w.color} key={w.value}>{w.label}</Tag>
                );
              }
              return;
          })}
        </Space>
      )
    },
    {
      title    : '更新时间',
      search   : false,
      dataIndex: 'update_time',
      render   : (_, record) => dayjs.unix(record?.update_time || 0).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      title    : '操作',
      dataIndex: 'option',
      valueType: 'option',
      width    : '20%',
      render   : (_, record) => (
        <>
          <Button
            type="primary"
            onClick={async () => {
              const res = await recipe(record.id || 0)
              if (res.code !== 0) {
                message.error(res.msg);
                return;
              }
              setRow(res.data)
              handleFormVisible(true)
            }}
          >
            编辑
          </Button>
          <Divider type="vertical"/>
          <Popconfirm
            title={`确认删除${record.title}吗?`}
            okText='是'
            cancelText='否'
            onConfirm={async () => {
              const res = await deleteRecipe(record.id || 0)
              if (res.code !== 0) {
                message.error(res.msg);
                return;
              }
              message.success("删除成功")
              actionRef.current?.reload();
            }
            }
          >
            <Button danger> 删除 </Button>
          </Popconfirm>
        </>
      )
    }
  ];

  return (
    <>
      <ProTable<API.RecipeItem>
        rowKey="id"
        headerTitle="菜谱管理"
        columns={columns}
        actionRef={actionRef}
        request={(params) => recipes(params.current || 0, params.pageSize || 0, {
          id         : params?.id || 0,
          category_id: params?.category_id || 0,
          title      : params?.title,
        })}
        toolBarRender={() => [
          <Button
            type="primary"
            onClick={async () => {
              setRow({});
              handleFormVisible(true);
            }}
          >
            新建菜谱
          </Button>
        ]}
      >
      </ProTable>

      {formVisible && (
        <RecipeForm
          onFinish={async (value) => {
            const res = await saveRecipe(value)
            if (res.code !== 0) {
              message.error(res.msg);
              return;
            }
            message.success("保存成功");
            handleFormVisible(false);
            actionRef.current?.reload();
          }}
          categoriesTreeData={categoriesTreeData || []}
          values={row || {}}
          onCancel={() => {
            handleFormVisible(false)
          }}
          formVisible={formVisible}
        >
        </RecipeForm>
      )}
    </>
  );
}

export default RecipeList;
