<template>
  <span class="headerAvatar">
    <!-- 头像 ，userStore.userInfo.headerImg-->
    <template v-if="picType === 'avatar'">
      <el-avatar v-if="userStore.userInfo.headerImg" :size="30" :src="avatar" />
      <el-avatar v-else :size="30" :src="noAvatar" />
    </template>
    <!-- 预览图片，props.picSrc-->
    <template v-if="picType === 'img'">
      <img v-if="props.picSrc" :src="file" class="avatar" />
      <img v-else :src="noAvatar" class="avatar" />
    </template>
    <!-- 文件,props.picSrc -->
    <template v-if="picType === 'file'">
      <el-image
        :src="file"
        class="file"
        :preview-src-list="previewSrcList"
        :preview-teleported="true"
      />
    </template>
  </span>
</template>

<script setup>
import noAvatarPng from '@/assets/noBody.png'
import zippic from '@/assets/zippic.png'
import { useUserStore } from '@/pinia/modules/user'
import { computed, ref } from 'vue'

defineOptions({
  name: 'CustomPic'
})

const props = defineProps({
  picType: {
    type: String,
    required: false,
    default: 'avatar'
  },
  picSrc: {
    type: String,
    required: false,
    default: ''
  },
  preview: {
    type: Boolean,
    default: false
  }
})

const path = ref(import.meta.env.VITE_BASE_API + '/')
const noAvatar = ref(noAvatarPng)
const zippic0 = ref(zippic)

const userStore = useUserStore()

const avatar = computed(() => {
  if (props.picSrc === '') {
    if (
      userStore.userInfo.headerImg !== '' &&
      userStore.userInfo.headerImg.slice(0, 4) === 'http'
    ) {
      return userStore.userInfo.headerImg
    }
    return path.value + userStore.userInfo.headerImg
  } else {
    if (props.picSrc !== '' && props.picSrc.slice(0, 4) === 'http') {
      return props.picSrc
    }
    return path.value + props.picSrc
  }
})
const file = computed(() => {
  //检查 props.picSrc 是否存在，以及它的前四个字符是否不等于 'http'
  if (props.picSrc && props.picSrc.slice(0, 4) !== 'http') {
    // 检查文件扩展名是否为 .zip
    if (props.picSrc.slice(-4) === '.zip') {
      return zippic0.value // 如果是 .zip 文件，返回默认图片路径
    }
    return path.value + props.picSrc //不是http开头的加上 path
  }
  return props.picSrc //http 开头的直接返回
})
// 预览图片，取决于 props.preview 和 file.value，只有在文件类型为图片时才会开启
const previewSrcList = computed(() => (props.preview ? [file.value] : []))
</script>

<style scoped>
.headerAvatar {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-right: 8px;
}
.file {
  width: 80px;
  height: 80px;
  position: relative;
}
</style>
