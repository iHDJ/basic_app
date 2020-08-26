import React from 'react';
import { connect } from 'dva';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { FormattedMessage, formatMessage } from 'umi';

@connect()
class ListPage extends React.Component {


  render() {


    return (
      <PageHeaderWrapper content={<FormattedMessage id="formandbasic-form.basic.description" />}>

      </PageHeaderWrapper>
    )
  }
}

export default ListPage
