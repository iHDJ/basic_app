export default [
  {
    name: "system_config",
    icon: "table",
    path: "/system_config",
    routes: [
      {
        name: "enumeration_list",
        path: "/system_config/enumerations",
        component: "./SystemConfig/Enumeration/ListPage",
      },
      {
        hideInMenu:true,
        name: "new_enumeration",
        path: "/system_config/enumerations/new",
        component: "./SystemConfig/Enumeration/NewPage",
      },
      {
        hideInMenu:true,
        name: "edit_enumeration",
        path: "/system_config/enumerations/:id/edit",
        component: "./SystemConfig/Enumeration/EditPage",
      },

      {
        name: "setting",
        path: "/system_config/setting",
        component: "./SystemConfig/SettingPage",
      },
    ]
  },
]
