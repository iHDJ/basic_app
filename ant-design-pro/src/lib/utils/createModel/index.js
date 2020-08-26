
export default function createModel(options) {
  console.log("createModel", options)
  const { namespace, state = {}, effects = {}, reducers = {} } = options;

  return {
    namespace,
    state,
    effects: {
      ...effects,
    },
    reducers: {
      ...reducers,
    },

    // reducers: parseReducers(
    //   {
    //     sync: syncReducer,
    //     update: updateReducer,
    //     ...reducers,
    //   },
    //   state
    // ),
  };
}
