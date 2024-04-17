<template>
  <el-row class="login-page">
    <el-col :span="12" class="bg"></el-col>
    <el-col :span="6" :offset="3" class="form">
      <!-- 注册相关表单 -->
      <el-form
        :model="formModel"
        :rules="rules"
        ref="form"
        size="large"
        autocomplete="off"
        v-if="isRegister"
      >
        <el-form-item>
          <h1>注册</h1>
        </el-form-item>
        <!-- <el-form-item prop="username">
          <el-input
            v-model="formModel.username"
            :prefix-icon="User"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item prop="repassword">
          <el-input
            v-model="formModel.repassword"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入再次密码"
          ></el-input>
        </el-form-item> -->
        <!-- <el-form-item>
          <el-button
            @click="register"
            class="button"
            type="primary"
            auto-insert-space
          >
            注册
          </el-button>
        </el-form-item>
        <el-form-item class="flex">
          <el-link type="info" :underline="false" @click="isRegister = false">
            ← 返回
          </el-link>
        </el-form-item> -->
      </el-form>

      <!-- 登录相关表单 -->
      <!-- @keyup.enter="submitForm" -->
      <el-form
        :model="loginFormData"
        :rules="rules"
        ref="loginForm"
        size="large"
        autocomplete="off"
        validate-on-rule-change="false"
        @keyup.enter="submitForm"
        v-else
      >
        <el-form-item>
          <h1>登录</h1>
        </el-form-item>
        <el-form-item prop="username">
          <el-input
            v-model="loginFormData.username"
            :prefix-icon="User"
            size="large"
            placeholder="请输入用户名"
            suffix-icon="user"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginFormData.password"
            :prefix-icon="Lock"
            show-password
            size="large"
            type="password"
            placeholder="请输入密码"
          />
        </el-form-item>
        <el-form-item v-if="loginFormData.openCaptcha" prop="captcha">
          <div>
            <el-input
              v-model="loginFormData.captcha"
              placeholder="请输入验证码"
              size="large"
            />
            <div>
              <img
                v-if="picPath"
                :src="picPath"
                alt="请输入验证码"
                @click="loginVerify()"
              />
            </div>
          </div>
        </el-form-item>

        <el-form-item class="flex">
          <div class="flex">
            <el-checkbox>记住我</el-checkbox>
            <el-link type="primary" :underline="false">忘记密码？</el-link>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button
            @click="login"
            class="button"
            type="primary"
            auto-insert-space
            >登录</el-button
          >
        </el-form-item>
        <el-form-item class="flex">
          <el-link type="info" :underline="false" @click="isRegister = true">
            注册 →
          </el-link>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<script setup>
import { captcha } from '@/api/user'
// import { userRegisterService, userLoginService } from '@/api/user.js'
import { User, Lock } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { ref, watch, reactive } from 'vue'
import { useUserStore } from '@/pinia'
import { useRouter } from 'vue-router'
const isRegister = ref(false)
const form = ref()

defineOptions({
  //设置组件的 name 选项为 'Login'。
  name: 'Login'
})

const router = useRouter()
// 验证函数
const checkUsername = (rule, value, callback) => {
  callback()
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6 && value.length != 0) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = () => {
  captcha({}).then(async (ele) => {
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    console.log(ele)
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  })
}
loginVerify() //当路由导航到这个页面时，立即执行

// 登录相关操作
const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'admin',
  password: '123456',
  captcha: '',
  captchaId: '',
  openCaptcha: false
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur'
    }
  ]
})

const userStore = useUserStore()
const login = async () => {
  return await userStore.LoginIn(loginFormData)
}
const submitForm = () => {
  loginForm.value.validate(async (v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true
      })
      loginVerify()
      return false
    }
  })
}

// // 整个的用于提交的form数据对象
// const formModel_register = ref({
//   username: '',
//   password: '',
//   repassword: ''
// })
// // 整个表单的校验规则
// const rules_register = {
//   username: [
//     { required: true, message: '请输入用户名', trigger: 'blur' },
//     { min: 5, max: 10, message: '用户名必须是5-10位的字符', trigger: 'blur' }
//   ],
//   password: [
//     { required: true, message: '请输入密码', trigger: 'blur' },
//     {
//       pattern: /^\S{6,15}$/,
//       message: '密码必须是6-15位的非空字符',
//       trigger: 'blur'
//     }
//   ],
//   repassword: [
//     { required: true, message: '请再次输入密码', trigger: 'blur' },
//     {
//       pattern: /^\S{6,15}$/,
//       message: '密码必须是6-15的非空字符',
//       trigger: 'blur'
//     },
//     {
//       validator: (rule, value, callback) => {
//         if (value !== formModel_register.value.password) {
//           callback(new Error('两次输入密码不一致!'))
//         } else {
//           callback()
//         }
//       },
//       trigger: 'blur'
//     }
//   ]
// }

// const register = async () => {
//   await form.value.validate()
//   //   console.log('开始注册请求')
//   await userRegisterService(formModel_register.value)
//   ElMessage.success('注册成功')
//   // 切换到登录
//   isRegister.value = false
// }
// //login，并存入pinia 和 持久化本地storage
// const userStore = useUserStore()
// // const router = useRouter()
// const login = async () => {
//   await form.value.validate()

//   // const res = await userLoginService(formModel.value) //被拒绝，await 会抛出一个错误
//   // userStore.setToken(res.data.token) //存入pinia
//   userStore.setToken('123456') //模拟登录成功，实际暂不发送请求

//   ElMessage.success('登录成功')
//   router.push('/')
// }

// // 切换的时候，重置表单内容
// watch(isRegister, () => {
//   formModel.value = {
//     username: '',
//     password: '',
//     repassword: ''
//   }
// })
</script>

<style lang="scss" scoped>
.login-page {
  height: 100vh;
  background-color: #fff;
  .bg {
    background:
      url('@/assets/logo2.png') no-repeat 60% center / 240px auto,
      url('@/assets/login_bg.jpg') no-repeat center / cover;
    border-radius: 0 20px 20px 0;
  }
  .form {
    display: flex;
    flex-direction: column;
    justify-content: center;
    user-select: none;
    .title {
      margin: 0 auto;
    }
    .button {
      width: 100%;
    }
    .flex {
      width: 100%;
      display: flex;
      justify-content: space-between;
    }
  }
}
</style>
