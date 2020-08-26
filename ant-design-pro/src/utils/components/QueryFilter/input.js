import React from 'react';
import { Input as AntInput, Form } from 'antd';

class Input extends React.Component {
  render() {
    const { name, label, required } = this.props;

    return (
      <React.Fragment>
        <Form.Item
          name={name}
          label={label}
        >
          <AntInput />
        </Form.Item>
      </React.Fragment>
    )
  }

}

export default Input;
