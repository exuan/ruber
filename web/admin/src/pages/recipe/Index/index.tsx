import React, {useEffect, useState} from 'react';
import ProForm, {
  ProFormText,
  ProFormUploadButton,
} from '@ant-design/pro-form';
import {imageRequest} from "@/utils/common";
import {Card, message} from "antd";
import {UploadFile} from "antd/lib/upload/interface";
import {recipeIndex, saveRecipeIndex} from '@/services/recipe';

const RecipeIndex: React.FC = () => {
  const [form] = ProForm.useForm();
  const [carousel, setCarousel] = useState<UploadFile[]>();

//  let carousel: UploadFile[] = [];

  useEffect(() => {
    (async () => {
      const res = await recipeIndex();
      if (!res || res.code !== 0) {
        message.error(res?.msg || "接口错误")
      }
      setCarousel( res?.data?.carousel ?
        res.data.carousel.map((v) => {
          return {
            uid     : v,
            status  : 'done',
            name    : v,
            url     : v,
            response: v
          }
        }) : []);
    })()

  }, []);

  return (
    <Card
      title="首页配置"
    >
      <ProForm
        layout="horizontal"
        labelCol={{span: 3}}
        form={form}
        onFinish={async (v) => {
          const res = await saveRecipeIndex(v)
          if (!res || res.code !== 0) {
            message.error(res?.msg ||"接口错误");
          } else {
            message.success("保存成功");
          }
        }}
        onReset={() => {
          form.resetFields();

        }}
      >

        <ProFormUploadButton
          fileList={carousel}
          label="轮播图"
          accept="image/*"
          fieldProps={{
            maxCount:5,
            customRequest: imageRequest
          }}
          onChange={(info) => {
            setCarousel(info.fileList);
            form.setFieldsValue({carousel: info.fileList.map((f) => (f.url ? f.url : f.response))});
            if (info.file.status === 'error') {
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

        <ProFormText hidden name="carousel"/>
      </ProForm>
    </Card>
  )
}

export default RecipeIndex;
