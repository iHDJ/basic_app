import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { connect } from 'dva';
import ProTable from '@ant-design/pro-table';
import { Card, Input, Select, Rate, Button } from 'antd';
import QueryFilter from "@/utils/components/QueryFilter"


@connect()
class EnumerationListPage extends React.Component {
  onSearch = (values) => {
    console.log(values)
  }

  render() {
    const columns = [
      {
        title: '枚举名',
        dataIndex: 'enumeration_name',
      },
      {
        title: "枚举Key",
        dataIndex: "key",
      },
      {
        title: "枚举值",
        dataIndex: "value",
      },
      {
        title: "创建时间",
        dataIndex: "created_at",
      },
      {
        title: "上次更新",
        dataIndex: "updated_at",
      },
    ]

    return (
      <PageHeaderWrapper>
        <Card bodyStyle={{padding: 15}}>
          <QueryFilter.Form
            hideReset={false}
            onSearch={this.onSearch}
          >
            <Input
              name="query"
              label="搜索"
              placeholder="搜索关键字"
            />

            <Select
              name="status"
              label="枚举状态"
              allowClear
            >
              <Select.Option value="is_system">系统的</Select.Option>
              <Select.Option value="is_opened">公开的</Select.Option>
              <Select.Option value="disable">不可用</Select.Option>
            </Select>

          </QueryFilter.Form>

          <Button type="primary">新建</Button>

          <ProTable search={false} columns={columns}>
            12333213131
          </ProTable>
        </Card>
      </PageHeaderWrapper>
    )
  }
}

export default EnumerationListPage;
