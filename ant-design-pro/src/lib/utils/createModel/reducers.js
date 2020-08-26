import _ from 'lodash';

export function syncReducer(state, { payload }) {
  const { response = {}, ...others } = payload;

  return {
    ...state,
    ...response,
    ...others,
  };
}

export function updateReducer(state, { payload }) {
  const { path, value } = payload;
  const nextState = _.cloneDeep(state);

  if (_.isPlainObject(value)) {
    const origValue = _.get(nextState, path);

    _.set(nextState, path, {
      ...origValue,
      ...value,
    });
  } else {
    _.set(nextState, path, value);
  }

  return nextState;
}

export function updateEntityReducer(state, { payload }) {
  const { type, id, attributes } = payload;

  const obj = _.get(state, ['entities', type, id]);

  return {
    ...state,
    entities: {
      ...state.entities,
      [type]: {
        ...state.entities[type],
        [id]: {
          ...obj,
          ...attributes,
        },
      },
    },
  };
}
