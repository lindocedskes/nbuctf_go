<template>
  <el-row class="login-page">
    <el-col :span="12" class="bg"></el-col>
    <el-col :span="6" :offset="3" class="form">
      <!-- 注册相关表单 -->
      <el-form
        :model="formModel_register"
        :rules="rules_register"
        ref="registerform"
        size="large"
        autocomplete="off"
        v-if="isRegister"
        class="space-y-4"
      >
        <el-form-item>
          <h1>注册</h1>
        </el-form-item>
        <el-form-item prop="username">
          <el-input
            v-model="formModel_register.username"
            :prefix-icon="User"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formModel_register.password"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item prop="repassword">
          <el-input
            v-model="formModel_register.repassword"
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入再次密码"
          ></el-input>
        </el-form-item>
        <el-form-item prop="nickName">
          <el-input
            v-model="formModel_register.nickName"
            :prefix-icon="User"
            placeholder="请输入昵称"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-input
            v-model="formModel_register.phone"
            :prefix-icon="Phone"
            placeholder="请输入联系手机号（选填）"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-input
            v-model="formModel_register.email"
            :prefix-icon="Message"
            placeholder="请输入联系邮箱（选填）"
          ></el-input>
        </el-form-item>
        <el-form-item
          v-if="formModel_register.openCaptcha"
          prop="captcha"
          class="flex items-center"
        >
          <img
            v-if="picPath"
            :src="picPath"
            alt="请输入验证码"
            @click="loginVerify()"
            class="w-40 h-20 ml-2"
          />
          <el-input
            v-model="formModel_register.captcha"
            placeholder="请输入验证码"
            size="large"
            class="flex-grow"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            @click="submitRegisterForm"
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
        </el-form-item>
      </el-form>

      <!-- 登录相关表单 -->
      <!-- @keyup.enter="submitForm" -->
      <el-form
        :model="loginFormData"
        :rules="rules"
        ref="loginForm"
        size="large"
        autocomplete="off"
        :validate-on-rule-change="false"
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
            @click="submitForm"
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
import { User, Lock, Message, Phone } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { ref, watch, reactive } from 'vue'
import { useUserStore } from '@/pinia'
const isRegister = ref(false)

defineOptions({
  //设置组件的 name 选项为 'Login'。
  name: 'Login'
})

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
    rules_register.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    console.log(ele)
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
    formModel_register.captchaId = ele.data.captchaId
    formModel_register.openCaptcha = ele.data.openCaptcha
  })
}
loginVerify() //当路由导航到这个页面时，立即执行

// 登录相关操作
const loginForm = ref(null) //与组件表单绑定，提供了一些方法，如 validate、resetFields
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

// 用于注册提交的form数据对象
const registerform = ref()
const formModel_register = reactive({
  username: '',
  password: '',
  repassword: '',
  nickName: '',
  phone: '',
  email: '',
  captcha: '',
  captchaId: '',
  openCaptcha: false
})
// 整个表单的校验规则
const rules_register = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, max: 10, message: '用户名必须是5-10位的字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    {
      pattern: /^\S{6,15}$/,
      message: '密码必须是6-15位的非空字符',
      trigger: 'blur'
    }
  ],
  nickName: [
    { required: true, message: '请输入昵称', trigger: 'blur' },
    { min: 2, max: 10, message: '昵称必须是2-10位的字符', trigger: 'blur' }
  ],
  repassword: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    {
      pattern: /^\S{6,15}$/,
      message: '密码必须是6-15的非空字符',
      trigger: 'blur'
    },
    {
      validator: (rule, value, callback) => {
        if (value !== formModel_register.password) {
          callback(new Error('两次输入密码不一致!'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur'
    }
  ]
}

import { registerbyuser } from '@/api/user'
const register = async () => {
  return await registerbyuser(formModel_register)
}

const submitRegisterForm = () => {
  registerform.value.validate(async (v) => {
    //校验表单
    if (v) {
      const flag = await register() //提交注册
      if (!flag) {
        loginVerify() // 刷新验证码
        ElMessage({
          type: 'error',
          message: '提交注册失败',
          showClose: true
        })
      } else {
        ElMessage({
          type: 'success',
          message: '注册成功，请登录',
          showClose: true
        })
        isRegister.value = false
        loginVerify() // 刷新验证码
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写注册信息',
        showClose: true
      })
      loginVerify()
      return false
    }
  })
}

// 切换的时候，重置表单内容
watch(isRegister, () => {
  // formModel_register.value = {
  //   username: '',
  //   password: '',
  //   repassword: ''
  // }
})
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
