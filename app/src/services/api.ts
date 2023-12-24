export const BurningApis = {
  auth: {
    login: {
      method: "POST",
      url: "/usercenter/api/v1/login",
      authenticated: false,
    },
    logout: {
      method: "POST",
      url: "/usercenter/api/v1/logout",
      authenticated: true,
    },
    register: {
      method: "POST",
      url: "/usercenter/api/v1/register",
      authenticated: false,
    },
    getProfile:{
      method: "GET",
      url: "/burning/api/v1/users/profile",
      authenticated: true,
    },
    updateProfile:{
      method: "PUT",
      url: "/burning/api/v1/users/profile",
      authenticated: true,
    },
    sign:{
      method: "POST",
      url: "/burning/api/v1/users/sign",
      authenticated: true,
    },
    getLastSign:{ //获取最近一次签到
      method: "GET",
      url: "/burning/api/v1/users/lastSign",
      authenticated: true,
    },
    getBodyInfo:{ //获取身体信息
      method: "GET",
      url: "/burning/api/v1/users/bodyInfo",
      authenticated: true,
    },
    updateBodyInfo:{ //更新身体信息
      method: "PUT",
      url: "/burning/api/v1/users/bodyInfo",
      authenticated: true,
    },
    createBodyInfo:{ //创建身体信息}
      method: "POST",
      url: "/burning/api/v1/users/bodyInfo",
      authenticated: true,
    },
  },
  home: {
    getVideoInfo: {
      method: "GET",
      url: "/burning/api/v1/video",
      authenticated: false,
    },
    getRandomMusic: {
      method: "GET",
      url: "/burning/api/v1/music",
      authenticated: false,
    },
  },
  history: {
    getHistory: {
      method: "GET",
      url: "/burning/api/v1/history/train",
      authenticated: true,
    },
    addHistory: {
      method: "POST",
      url: "/burning/api/v1/history/train",
      authenticated: true,
    },
    deleteHistory: {
      method: "DELETE",
      url: (trainID: number) => {
        return `/burning/api/v1/history/train/${trainID}`;
      },
      authenticated: true,
    },
    updateHistory: {
      method: "PUT",
      url: (trainID: number) => {
        return `/burning/api/v1/history/train/${trainID}`;
      },
      authenticated: true,
    },
    finishTrain:{
      method: "POST",
      url: (trainID: number) => {
        return `/burning/api/v1/history/train/${trainID}/finish`;
      },
      authenticated: true,
    },
    getTrainHistoryDetail: {
      method: "GET",
      url: (trainID: number) => {
        return `/burning/api/v1/history/train/${trainID}`;
      },
      authenticated: true,
    },
    addNewTrainContent: {
      method: "POST",
      url: (trainID: number) => {
        return `/burning/api/v1/history/train/${trainID}/content`;
      },
      authenticated: true,
    },
    deleteTrainContent: {
      method: "DELETE",
      url: (trainID: number, contentID: number) => {
        return `/burning/api/v1/history/train/${trainID}/content/${contentID}`;
      },
      authenticated: true,
    },
    updateTrainContent: {
      method: "PUT",
      url: (trainID: number, contentID: number) => {
        return `/burning/api/v1/history/train/${trainID}/content/${contentID}`;
      },
      authenticated: true,
    },
    //饮食记录
    getFoodHistory: {
      method: "GET",
      url: "/burning/api/v1/history/food",
      authenticated: true,
    },
    updateFoodHistory: {
      method: "PUT",
      url: (foodID: number) => {
        return `/burning/api/v1/history/food/${foodID}`;
      },
      authenticated: true,
    },
  },
  actions:{
    getActions:{
      method: "GET",
      url: "/burning/api/v1/actions",
      authenticated: false,
    },
    addAction:{
      method: "POST",
      url: "/burning/api/v1/actions",
      authenticated: false,
    },
    deleteAction:{
      method: "DELETE",
      url: (actionID: number) => {
        return `/burning/api/v1/actions/${actionID}`;
      },
      authenticated: false,
    },
    updateAction:{
      method: "PUT",
      url: (actionID: number) => {
        return `/burning/api/v1/actions/${actionID}`;
      },
      authenticated: false,
    },
    addCustomAction:{
      method: "POST",
      url: "/burning/api/v1/actions/custom",
      authenticated: true,
    },
    getCustomActions:{
      method: "GET",
      url: "/burning/api/v1/actions/custom",
      authenticated: true,
    },
    deleteCustomAction:{
      method: "DELETE",
      url: (actionID: number) => {
        return `/burning/api/v1/actions/custom/${actionID}`;
      },
      authenticated: true,
    },
    updateCustomAction:{
      method: "PUT",
      url: (actionID: number) => {
        return `/burning/api/v1/actions/custom/${actionID}`;
      },
      authenticated: true,
    },
  },
};
