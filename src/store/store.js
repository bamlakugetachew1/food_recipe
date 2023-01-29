import Vuex from 'vuex'
export default new Vuex.Store({
    state:{
      isitemdeletd:0,
      },
      getters: {
        getMessage(state){
          return state.isitemdeletd
        },
      },
      mutations:{
        changeMessageValue(state, message){
          state.isitemdeletd = message
        },
      }
});
