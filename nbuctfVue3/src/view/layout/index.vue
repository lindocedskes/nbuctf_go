<template>
  <el-container class="layout-container">
    <el-container>
      <el-header>
        <div class="el-header__logo"></div>
        <el-menu
          mode="horizontal"
          active-text-color="#ffd04b"
          background-color="#232323"
          :default-active="$route.path"
          text-color="#fff"
          router
        >
          <el-menu-item index="/layout/dashboard">
            <el-icon><Management /></el-icon>
            <span>仪表盘</span>
          </el-menu-item>
          <el-menu-item index="/layout/announcement">
            <el-icon><Promotion /></el-icon>
            <span>公告</span>
          </el-menu-item>
          <el-menu-item index="/article/manage">
            <el-icon><Promotion /></el-icon>
            <span>比赛</span>
          </el-menu-item>
          <el-menu-item index="/article/manage">
            <el-icon><Promotion /></el-icon>
            <span>排行榜</span>
          </el-menu-item>
          <el-menu-item index="/layout/about">
            <el-icon><Promotion /></el-icon>
            <span>关于</span>
          </el-menu-item>
          <el-menu-item index="/layout/announcement_manage">
            <el-icon><Promotion /></el-icon>
            <span>公告管理</span>
          </el-menu-item>
          <el-menu-item index="/layout/userbyadmin">
            <el-icon><Promotion /></el-icon>
            <span>用户管理</span>
          </el-menu-item>
        </el-menu>
        <el-dropdown placement="bottom-end" @command="handleCommand">
          <!-- 展示给用户，默认看到的 -->
          <span class="el-dropdown__box">
            <el-avatar :src="userStore.userInfo.headerImg || avatar" />
            <el-icon><CaretBottom /></el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <div>
                user：<strong>{{
                  userStore.userInfo.nickName || userStore.userInfo.userName
                }}</strong>
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
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
      <el-footer></el-footer>
    </el-container>
  </el-container>
</template>

<script setup>
import {
  Management,
  Promotion,
  User,
  SwitchButton,
  CaretBottom
} from '@element-plus/icons-vue'
import avatar from '@/assets/default_avatar.png'

import { useUserStore } from '@/pinia'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const router = useRouter()

onMounted(() => {
  // userStore.getUser() // 每次请求permisssion都会获取
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
  .el-header {
    &__logo {
      //选择为：header__logo
      width: 60px;
      height: 60px;
      background: url('@/assets/app_logo.png') no-repeat center / 120px auto;
    }
    .el-menu {
      flex: 1;
      display: flex;
      justify-content: flex-end;
      gap: 20px; // 根据需要调整间隙
    }
  }
  .el-header {
    background-color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;
    .el-dropdown__box {
      display: flex;
      align-items: center;
      .el-icon {
        color: #999;
        margin-left: 10px;
      }

      &:active,
      &:focus {
        outline: none;
      }
    }
  }
  .el-footer {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    color: #666;
  }
}
</style>
