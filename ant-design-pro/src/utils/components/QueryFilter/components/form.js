import React from 'react';
import { Form, Input, Row, Col, Button } from 'antd';

class QueryForm extends React.Component {

  onSearch = (values) => {
    console.log(values)
  }

  componentDidMount() {
    const url = new URL(location.href)
    React.Children.forEach(this.props.children, children => {
      const { name } = children.props;
      const urlQueryValue = url.searchParams.get(name)
      console.log("query", name, "value", urlQueryValue)

      this.form.setFieldsValue({ [name]: urlQueryValue })
    })
  }

  render() {
    const { children, hideReset } = this.props;

    return (
      <React.Fragment>
        <Form
          ref={form => this.form = form}
          name="search"
          onFinish={this.onSearch}
        >
          <Row gutter={24}>
            {
              React.Children.map(children, (children) => {
                console.log(children)
                return (
                  <Col lg={6} xs={24}>
                    <Form.Item
                      name={children.props.name}
                      label={children.props.label}
                    >
                      {children}
                    </Form.Item>
                  </Col>
                )
              })
            }


          </Row>

          <Col span={24} style={{ textAlign: 'right' }}>
            <Button type="primary" htmlType="submit">
              搜索
            </Button>

            {!hideReset && (
              <Button
                style={{ margin: '0 8px' }}
                onClick={() => {
                  this.form.resetFields();
                }}
              >
                重置
              </Button>
            )}
            {/* <a
              style={{ fontSize: 12 }}
              onClick={() => {
                setExpand(!expand);
              }}
            >
              {expand ? <UpOutlined /> : <DownOutlined />} Collapse
            </a> */}
          </Col>
        </Form>
      </React.Fragment>
    )
  }
}

export default QueryForm;
