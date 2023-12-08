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
    }
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
  },

};
