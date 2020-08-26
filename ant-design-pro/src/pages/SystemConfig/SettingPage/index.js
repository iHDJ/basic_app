import React from 'react';
import { connect } from 'dva';
import { formatMessage } from 'umi';
import { Card, Input, Form } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';

const formItemLayout = {
  labelCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 7,
    },
  },
  wrapperCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 12,
    },
    md: {
      span: 10,
    },
  },
};

@connect()
class SystemConfigSettingPage extends React.Component {
  //formRef = React.createRef();

  render() {

    return (
      <PageHeaderWrapper content={formatMessage({id: "system_config.setting.description"})}>
        <Card bordered={false}>
          <Form
            style={{
              marginTop: 8,
            }}
            //ref={this.formRef}
            //onFinish={onFinish}
            //onFinishFailed={onFinishFailed}
            //onValuesChange={onValuesChange}
          >
            <Form.Item
            {...formItemLayout}
              label="域名"
              name="domain"
              rules={[
                {
                  required: true,
                  message: '请输入域名',
                },
              ]}
            >
              <Input />
            </Form.Item>
          </Form>
        </Card>
      </PageHeaderWrapper>
    )
  }
}

export default SystemConfigSettingPage;
