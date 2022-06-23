import React, {useState} from 'react';
import ProForm, {ModalForm, ProFormDigit, ProFormSelect, ProFormText, ProFormUploadButton} from '@ant-design/pro-form';

import {RecipeWeight} from '@/utils/constant';
import {message, TreeSelect} from 'antd';
import {imageRequest} from "@/utils/common";
import type {DataNode} from "rc-tree-select/lib/interface";
import type {UploadFile} from "antd/lib/upload/interface";
import ReactWEditor from 'wangeditor-for-react';
import {result} from "lodash";

export type RecipeFormType = {
  id?: number;
  category_id?: number;
  title?: string;
  image?: string;
  content?: string;
  weight?: number;
  sort?: number;
} & Partial<API.RecipeItem>;

export type RecipeFormProps = {
  onCancel: () => void;
  onFinish: (values: RecipeFormType) => Promise<void>;
  formVisible: boolean;
  values: Partial<API.RecipeItem>;
  categoriesTreeData: DataNode[];
}

const RecipeForm: React.FC<RecipeFormProps> = (props) => {
  const [form] = ProForm.useForm();

  const {onCancel, onFinish, categoriesTreeData, formVisible, values} = props;
  const [imageList, setImageList]                                     = useState<UploadFile[]>(values.image ? [{
    uid     : values.image,
    status  : 'done',
    name    : values.image,
    url     : values.image,
    response: values.image
  }] : []);

  const [weightSelect, setWeightSelect] = useState<number[]>(() => {
    const selects: number[] = [];
    [RecipeWeight.Hot, RecipeWeight.Favorite].forEach((v) => {
      if (((values.weight || 0) & v) === v) {
        selects.push(v)
      }
    })
    return selects;
  });


  return (
    <ModalForm
      width={1000}
      layout="horizontal"
      labelCol={{span: 2}}
      title={values.id ? `编辑 ${[values.title]}` : `新建`}
      initialValues={values}
      form={form}
      visible={formVisible}
      modalProps={{
        onCancel: () => {
          form.resetFields();
          onCancel();
        }
      }}
      onFinish={onFinish}

      submitter={{
        searchConfig: {
          submitText: "保存",
          resetText : "关闭"
        }
      }}
    >
      <ProFormText
        name="title"
        label="名称"
        placeholder="请输入菜谱名称"
        rules={[
          {
            required: true,
            message : '请输入菜谱名称！',
          },
        ]}
      />

      <ProForm.Item
        label="分类"
        name="category_id"
        rules={[
          {
            required: true,
            message : '请选择分类！',
          },
        ]}
      >
        <TreeSelect

          treeData={categoriesTreeData}
          placeholder="请选择分类"
        >
        </TreeSelect>
      </ProForm.Item>

      <ProFormDigit
        label="排序"
        name="sort"
        placeholder="数字越大越靠前"
        min={0}
        max={100000000}
        fieldProps={{precision: 0}}
      />

      <ProFormSelect
        name="weight_select"
        label="权重"
        fieldProps={{
          value   : weightSelect,
          mode    : 'multiple',
          onChange: (value) => {
            let weight = 0;
            setWeightSelect(value);
            value.forEach((v: number) => weight += v)
            form.setFieldsValue({weight: weight});
          },
        }}
        options={[
          {
            value: RecipeWeight.Hot,
            label: "热门",
          },
          // {
          //   value: RecipeWeight.Favorite,
          //   label: "最爱",
          // },
        ]}
      />

      <ProFormUploadButton
        name="image"
        label="图片"
        fileList={imageList}
        accept="image/*"
        max={1}
        fieldProps={{
          customRequest: imageRequest
        }}
        onChange={(info) => {
          setImageList(info.fileList);
          if (info.file.status === 'removed') {
            form.setFieldsValue({image: ''});
          } else if (info.file.status === 'done') {
            form.setFieldsValue({image: info.file.response});
          } else if (info.file.status === 'error') {
            message.error(`${info.file.name} 上传失败`);
          }
        }}
        rules={[
          {
            required: true,
            message : '请上传图片！',
          },
        ]}
      />

      <ProForm.Item
        label="内容"
        name="content"
        rules={[
          {
            required: true,
            message : '请输入内容',
          },
        ]}
      >
        <ReactWEditor
          config={{
            uploadImgMaxLength: 1,
            customUploadImg   : async (resultFiles: any[], insertImgFn: any) => {
              await imageRequest({file: resultFiles[0], onSuccess: (u, result) => {
                  if (u.length > 0) {
                    insertImgFn(u)
                  }
                }
              })
            },
          }}
          defaultValue={values.content}
          onChange={(html) => {
            form.setFieldsValue({content: html});
          }}
        />

      </ProForm.Item>

      <ProFormText hidden name="content"/>
      <ProFormText hidden name="weight"/>
      <ProFormText hidden name="image"/>
      <ProFormText hidden name="id"/>
    </ModalForm>
  )
}
export default RecipeForm;
