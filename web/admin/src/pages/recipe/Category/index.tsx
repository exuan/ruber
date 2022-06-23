import React, {useState, useRef} from 'react';
import ProTable, {ProColumns, ActionType} from "@ant-design/pro-table";
import {Button, Divider, Popconfirm, message} from 'antd';
import {recipeCategories, saveRecipeCategory, recipeCategory, deleteRecipeCategory} from '@/services/recipe';
import RecipeCategoryForm from "@/pages/recipe/Category/components/RecipeCategoryForm";
import * as dayjs from 'dayjs'
import {tree} from '@/utils/common';

const RecipeCategoryList: React.FC = () => {
  const [formVisible, handleFormVisible] = useState<boolean>(false)
  const [row, setRow]                    = useState<API.RecipeCategoryItem>()
  const actionRef                        = useRef<ActionType>()

  const parentRequest = async () => {
    const res = await recipeCategories(0, 0, {all: 1, parent_id: "0", not_id: row?.id?.toString() || ""});
    if (!res.success) {
      message.error('获取失败');
      return [];
    }

    const categories = [{label: " 未选择 ", value: 0}];
    res.data.forEach((r) => {
      if ((r.parent_id || 0) !== 0) {
        return;
      }
      categories.push({label: r.name || "", value: r.id || 0});
    })
    return categories;
  };

  const columns: ProColumns<API.RecipeCategoryItem>[] = [
    {
      title    : 'ID',
      dataIndex: 'id',
      width    : '20%',
    },
    {
      title    : '名称',
      dataIndex: 'name',
      width    : '20%',
    },
    {
      title    : 'parent_id',
      dataIndex: 'parent_id',
    },
    {
      title    : '排序',
      dataIndex: 'sort',
    },
    {
      title    : '更新时间',
      dataIndex: 'update_time',
      width    : '20%',
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
              const res = await recipeCategory(record.id || 0)
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
            title={`确认删除${record.name}吗?`}
            okText='是'
            cancelText='否'
            onConfirm={async () => {
              const res = await deleteRecipeCategory(record.id || 0)
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
  ]

  return (
    <>
      <ProTable<API.RecipeCategoryItem>
        rowKey="id"
        actionRef={actionRef}
        headerTitle="分类管理"
        search={false}
        columns={columns}
        request={async (params) => {
          const res = await recipeCategories(params.page || 0, params.size || 0, {
            all      : 1,
            parent_id: "",
            not_id   : ""
          })
          res.data  = tree(res.data, 0, 'parent_id')
          return res
        }}
        pagination={false}
        toolBarRender={() => [
          <Button
            type="primary"
            onClick={() => {
              setRow({});
              handleFormVisible(true)
            }}
          >
            新建分类
          </Button>
        ]}
      >
      </ProTable>

      {formVisible && (
        <RecipeCategoryForm
          onCancel={() => {
            handleFormVisible(false)
          }}
          onFinish={async (value) => {
            const res = await saveRecipeCategory(value)
            if (res.code !== 0) {
              message.error(res.msg);
              return;
            }
            message.success("保存成功");
            handleFormVisible(false);
            actionRef.current?.reload();
          }}

          formVisible={formVisible}
          parentRequest={parentRequest}
          values={row || {}}
        >
        </RecipeCategoryForm>
      )}
    </>
  )
}
export default RecipeCategoryList;
