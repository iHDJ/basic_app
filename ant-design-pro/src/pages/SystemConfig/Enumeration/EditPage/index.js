import React from 'react';
import { connect } from 'dva';

@connect()
class EnumerationEditPage extends React.Component {
  render() {
    return (
      <PageHeaderWrapper content="xx">
        <div>this is enumeration edit page</div>
      </PageHeaderWrapper>
    )
  }
}

export default EnumerationEditPage;
