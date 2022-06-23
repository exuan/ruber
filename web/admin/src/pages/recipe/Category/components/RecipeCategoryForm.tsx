import React from 'react';

import ProForm, {
  ModalForm,
  ProFormText,
  ProFormSelect,
  ProFormDigit,
} from '@ant-design/pro-form';


import {ProFieldRequestData} from "@ant-design/pro-utils/lib/typing";

export type  RecipeCategoryFormType = {
  id?: number;
  parent_id?: number;
  name?: string;
  sort?: number;

} & Partial<API.RecipeCategoryItem>

export type RecipeCategoryFormProps = {
  onCancel: () => void;
  onFinish: (values: RecipeCategoryFormType) => Promise<void>;
  formVisible: boolean;
  values: Partial<API.RecipeCategoryItem>
  parentRequest: ProFieldRequestData;
}

const RecipeCategoryForm: React.FC<RecipeCategoryFormProps> = (props) => {
  const [form]                                                   = ProForm.useForm();
  const {onCancel, onFinish, formVisible, parentRequest, values} = props;

  return (
    <ModalForm
      layout="horizontal"
      form={form}
      visible={formVisible}
      modalProps={{
        onCancel: (_) => {
          form.resetFields();
          onCancel();
        }
      }}
      labelCol={{span: 2}}
      title={props.values.id ? `编辑 ${[props.values.name]}` : `新建`}
      initialValues={values}
      submitter={{
        searchConfig: {
          submitText: "保存",
          resetText : "关闭"
        }
      }}
      onFinish={async (v) => {
        await onFinish(v)
      }}
    >
      <ProFormText
        name="name"
        label="名称"
        placeholder="请输入分类名称"
        rules={[
          {
            required: true,
            message : '请输入分类名称！',
          },
        ]}
      />
      <ProFormSelect
        name="parent_id"
        label="父级"
        placeholder="请选择父级"
        request={parentRequest}
      />

      <ProFormDigit
        label="排序"
        name="sort"
        placeholder="数字越大越靠前"
        min={0}
        max={100000000}
        fieldProps={{precision: 0}}
      />
      <ProFormText hidden name="id"/>
    </ModalForm>
  );
}

export default RecipeCategoryForm;
