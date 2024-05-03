<template>
  <nav class="bg-black border-gray-200 h-16 backgroundNav">
    <div
      class="max-w-screen-xl flex items-center justify-between mx-4 p-0 overflow-auto"
    >
      <div class="flex-grow-0">
        <a class="flex items-center space-x-2 rtl:space-x-reverse ml-0">
          <dv-decoration-11 style="width: 200px; height: 60px">
            <span
              class="self-center text-2xl font-semibold whitespace-nowrap text-green-300 font-semibold"
              >NBUCTF</span
            >
          </dv-decoration-11>
          <!-- <img src="@/assets/app_logo.png" class="h-12" alt="Flowbite Logo" /> -->
        </a>
      </div>
      <div class="flex-grow">
        <el-menu
          mode="horizontal"
          active-text-color="#ffd04b"
          background-color="#000"
          :default-active="$route.path"
          text-color="#fff"
          router
          class="w-full flex justify-end"
        >
          <el-menu-item index="/layout/gameintroduce">
            <el-icon><Promotion /></el-icon>
            <span>比赛介绍</span>
          </el-menu-item>
          <el-menu-item index="/layout/announcement">
            <el-icon><Promotion /></el-icon>
            <span>公告</span>
          </el-menu-item>
          <el-menu-item index="/layout/game">
            <el-icon><Promotion /></el-icon>
            <span>赛题</span>
          </el-menu-item>
          <el-menu-item index="/layout/rank">
            <el-icon><Promotion /></el-icon>
            <span>排行榜</span>
          </el-menu-item>
          <el-menu-item index="/layout/about">
            <el-icon><Promotion /></el-icon>
            <span>关于</span>
          </el-menu-item>
          <el-sub-menu index="2" v-if="userStore.userInfo.authorityId === 888">
            <template #title>管理员</template>
            <el-menu-item index="/layout/admin/gameadmin">
              <el-icon><Promotion /></el-icon>
              <span>比赛管理</span>
            </el-menu-item>
            <el-menu-item index="/layout/admin/user">
              <el-icon><Promotion /></el-icon>
              <span>用户管理</span>
            </el-menu-item>
            <el-menu-item index="/layout/file">
              <el-icon><Promotion /></el-icon>
              <span>文件管理</span>
            </el-menu-item>
          </el-sub-menu>
        </el-menu>
      </div>

      <div class="flex-grow-0">
        <el-dropdown placement="bottom-end" @command="handleCommand">
          <!-- 展示给用户，默认看到的 -->
          <span class="el-dropdown__box mr-4">
            <el-avatar
              :src="imgSrc"
              class="text-sm bg-gray-800 rounded-full md:me-0 focus:ring-4 focus:ring-gray-300"
            />
            <el-icon><CaretBottom /></el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <div class="px-4 py-3">
                <span class="block text-sm text-gray-900">{{
                  userStore.userInfo.nickName || userStore.userInfo.userName
                }}</span>
                <span class="block text-sm text-gray-500 truncate">{{
                  userStore.userInfo.email
                }}</span>
              </div>
              <el-dropdown-item command="person" :icon="User"
                >个人信息</el-dropdown-item
              >
              <el-dropdown-item command="logout" :icon="SwitchButton"
                >退出登录</el-dropdown-item
              >
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </nav>
  <el-container class="layout-container">
    <dv-border-box-1 style="height: 100%; width: 100%">
      <el-main>
        <router-view></router-view>
      </el-main>
    </dv-border-box-1>

    <el-footer></el-footer>
  </el-container>
</template>

<script setup>
import {
  Promotion,
  User,
  SwitchButton,
  CaretBottom
} from '@element-plus/icons-vue'
import noAvatarPng from '@/assets/default_avatar.png'

import { useUserStore } from '@/pinia'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { computed, ref } from 'vue'
// import { emitter } from '@/utils/bus.js'
//useRoute 返回的是当前的路由对象； useRouter 返回的是一个路由实例对象，跳转、获取所有的路由信息等
// import { useRoute } from 'vue-router'

const userStore = useUserStore()
const router = useRouter()

const path = ref(import.meta.env.VITE_BASE_API + '/') // 请求地址默认同源，/api 标志为代理地址，转发到后端
const noAvatar = ref(noAvatarPng)

const imgSrc = computed(() => {
  if (userStore.userInfo.headerImg !== '') {
    if (userStore.userInfo.headerImg.slice(0, 4) === 'http') {
      return userStore.userInfo.headerImg
    } else {
      return path.value + userStore.userInfo.headerImg
    }
  }
  return noAvatar.value //noAvatar 是一个 ref对象，.value 才是值
})

onMounted(() => {
  // userStore.getUser() // 每次请求permisssion都会获取
  // 挂载一些通用的事件
})

const handleCommand = async (key) => {
  if (key === 'logout') {
    // 退出操作
    await ElMessageBox.confirm('你确认要进行退出么', '温馨提示', {
      type: 'warning',
      confirmButtonText: '确认',
      cancelButtonText: '取消'
    })

    // 清除本地的数据 (token + user信息)
    userStore.LoginOut()
  } else {
    // 跳转操作
    router.push(`/layout/${key}`)
  }
}
</script>
<style lang="scss" scoped>
.layout-container {
  height: 100vh;
  background-image: url('/src/assets/login_bg.png');
  background-size: cover;
  background-position: center;
  overflow-y: auto; /* 当内容超过容器的高度时，使用滚动 */

  .el-menu {
    flex: 1;
    display: flex;
    justify-content: flex-end;
    gap: 20px; // 根据需要调整间隙
  }

  // .el-header {
  //   background-color: #fff;
  //   display: flex;
  //   align-items: center;
  //   justify-content: space-between;
  //   .el-dropdown__box {
  //     display: flex;
  //     align-items: center;
  //     .el-icon {
  //       color: #999;
  //       margin-left: 10px;
  //     }

  //     &:active,
  //     &:focus {
  //       outline: none;
  //     }
  //   }
  // }
  .el-footer {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    color: #666;
  }
}
.backgroundNav {
  background-image: url('/src/assets/intro_bg.png');
  background-size: 120% 100%;
  background-position: center;
}
.el-menu {
  background-color: transparent;
  border-bottom: 0;
}
.el-menu .el-menu-item {
  background-color: transparent !important;
  color: #80ffff !important;
}
.el-menu .el-menu-item.is-active {
  color: #e36ce9 !important;
  border-color: #e36ce9 !important;
}
.el-menu .el-menu-item {
  background-color: transparent !important;
  color: #80ffff !important;
}
</style>
<style lang="scss">
// 有二级菜单的标题及子菜单菜单字体颜色
.el-menu .el-sub-menu__title {
  color: #80ffff !important;
  border-bottom: 0;
}

.el-menu .el-submenu .el-menu {
  background-color: transparent !important;
}
// 有二级菜单的标题时鼠标悬停的颜色
.el-sub-menu__title.el-tooltip__trigger.el-tooltip__trigger:hover {
  color: red;
}
</style>
