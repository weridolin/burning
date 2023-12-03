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
      url: "/mini/burning/api/v1/profile",
      authenticated: true,
    },
    updateProfile:{
      method: "PUT",
      url: "/mini/burning/api/v1/profile",
      authenticated: true,
    },
  },
  home: {
    getVideoInfo: {
      method: "GET",
      url: "/mini/burning/api/v1/video",
      authenticated: false,
    },
    getRandomMusic: {
      method: "GET",
      url: "/mini/burning/api/v1/music",
      authenticated: false,
    },
  },
  history: {
    getHistory: {
      method: "GET",
      url: "/mini/burning/api/v1/history/train",
      authenticated: true,
    },
    addHistory: {
      method: "POST",
      url: "/mini/burning/api/v1/history/train",
      authenticated: true,
    },
    deleteHistory: {
      method: "DELETE",
      url: (trainID: number) => {
        return `/mini/burning/api/v1/history/train/${trainID}`;
      },
      authenticated: true,
    },
    updateHistory: {
      method: "PUT",
      url: (trainID: number) => {
        return `/mini/burning/api/v1/history/train/${trainID}`;
      },
      authenticated: true,
    },
    getTrainHistoryDetail: {
      method: "GET",
      url: (trainID: number) => {
        return `/mini/burning/api/v1/history/train/${trainID}`;
      },
      authenticated: true,
    },
    addNewTrainContent: {
      method: "POST",
      url: (trainID: number) => {
        return `/mini/burning/api/v1/history/train/${trainID}/content`;
      },
      authenticated: true,
    },
    deleteTrainContent: {
      method: "DELETE",
      url: (trainID: number, contentID: number) => {
        return `/mini/burning/api/v1/history/train/${trainID}/content/${contentID}`;
      },
      authenticated: true,
    },
    updateTrainContent: {
      method: "PUT",
      url: (trainID: number, contentID: number) => {
        return `/mini/burning/api/v1/history/train/${trainID}/content/${contentID}`;
      },
      authenticated: true,
    },
    //饮食记录
    getFoodHistory: {
      method: "GET",
      url: "/mini/burning/api/v1/history/food",
      authenticated: true,
    },
    updateFoodHistory: {
      method: "PUT",
      url: (foodID: number) => {
        return `/mini/burning/api/v1/history/food/${foodID}`;
      },
      authenticated: true,
    },
  },
  actions:{
    getActions:{
      method: "GET",
      url: "/mini/burning/api/v1/actions",
      authenticated: false,
    },
    addAction:{
      method: "POST",
      url: "/mini/burning/api/v1/actions",
      authenticated: false,
    },
    deleteAction:{
      method: "DELETE",
      url: (actionID: number) => {
        return `/mini/burning/api/v1/actions/${actionID}`;
      },
      authenticated: false,
    },
    updateAction:{
      method: "PUT",
      url: (actionID: number) => {
        return `/mini/burning/api/v1/actions/${actionID}`;
      },
      authenticated: false,
    },
  },

};
