import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { connect } from 'dva';

@connect()
class EnumerationNewPage extends React.Component {
  render() {

    return (
      <PageHeaderWrapper content="xx">
        <div>this is enumeration new page</div>
      </PageHeaderWrapper>
    )
  }
}

export default EnumerationNewPage;
