export const BurningApis = {
  auth: {
    login: {
      method: "post",
      url: "/usercenter/api/v1/login",
      authenticated: false,
    },
    logout: {
      method: "post",
      url: "/usercenter/api/v1/logout",
      authenticated: true,
    },
    register: {
      method: "post",
      url: "/usercenter/api/v1/register",
      authenticated: false,
    },
    getProfile:{
      method: "get",
      url: "/mini/burning/api/profile",
      authenticated: true,
    },
    updateProfile:{
      method: "put",
      url: "/mini/burning/api/profile",
      authenticated: true,
    },
  },
  home: {
    getVideoInfo: {
      method: "get",
      url: "/mini/burning/api/video",
      authenticated: false,
    },
    getRandomMusic: {
      method: "get",
      url: "/mini/burning/api/music",
      authenticated: false,
    },
  },
  history: {
    getHistory: {
      method: "get",
      url: "/mini/burning/api/history",
      authenticated: true,
    },
    addHistory: {
      method: "post",
      url: "/mini/burning/api/history/train",
      authenticated: true,
    },
    deleteHistory: {
      method: "delete",
      url: (trainID: number) => {
        return `/mini/burning/api/history/train/${trainID}`;
      },
      authenticated: true,
    },
    updateHistory: {
      method: "put",
      url: (trainID: number) => {
        return `/mini/burning/api/history/train/${trainID}`;
      },
      authenticated: true,
    },
    getTrainHistoryDetail: {
      method: "get",
      url: (trainID: number) => {
        return `/mini/burning/api/history/train/${trainID}`;
      },
      authenticated: true,
    },
    addNewTrainContent: {
      method: "post",
      url: (trainID: number) => {
        return `/mini/burning/api/history/train/${trainID}/content`;
      },
      authenticated: true,
    },
    deleteTrainContent: {
      method: "delete",
      url: (trainID: number, contentID: number) => {
        return `/mini/burning/api/history/train/${trainID}/content/${contentID}`;
      },
      authenticated: true,
    },
    updateTrainContent: {
      method: "put",
      url: (trainID: number, contentID: number) => {
        return `/mini/burning/api/history/train/${trainID}/content/${contentID}`;
      },
      authenticated: true,
    },
    //饮食记录
    getFoodHistory: {
      method: "get",
      url: "/mini/burning/api/history/food",
      authenticated: true,
    },
    updateFoodHistory: {
      method: "put",
      url: (foodID: number) => {
        return `/mini/burning/api/history/food/${foodID}`;
      },
      authenticated: true,
    },
  },
  actions:{
    getActions:{
      method: "get",
      url: "/mini/burning/api/actions",
      authenticated: false,
    },
    addAction:{
      method: "post",
      url: "/mini/burning/api/actions",
      authenticated: false,
    },
    deleteAction:{
      method: "delete",
      url: (actionID: number) => {
        return `/mini/burning/api/actions/${actionID}`;
      },
      authenticated: false,
    },
    updateAction:{
      method: "put",
      url: (actionID: number) => {
        return `/mini/burning/api/actions/${actionID}`;
      },
      authenticated: false,
    },
  },

};
