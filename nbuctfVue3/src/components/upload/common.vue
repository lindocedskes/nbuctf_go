<template>
  <div>
    <el-upload
      :action="`${path}/file/upload`"
      :before-upload="checkFile"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      class="upload-btn"
    >
      <el-button type="primary">普通上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { isVideoMime, isImageMime } from '@/utils/image'

defineOptions({
  name: 'UploadCommon'
})

const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isLt5M = file.size / 1024 / 1024 < 5 // 5MB
  const isLt50M = file.size / 1024 / 1024 < 50 // 50MB
  const isVideo = isVideoMime(file.type)
  const isImage = isImageMime(file.type)
  const isZip = file.type === 'application/zip'

  let pass = true
  if (!isVideo && !isImage && !isZip) {
    ElMessage.error(
      '上传文件只能是zip格式，上传图片只能是 jpg,png,svg,webp 格式, 上传视频只能是 mp4,webm 格式!'
    )
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt50M && isVideo) {
    ElMessage.error('上传视频大小不能超过 50MB')
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt5M && isImage) {
    ElMessage.error('未压缩的上传图片大小不能超过 5M，请使用压缩上传')
    fullscreenLoading.value = false
    pass = false
  }

  // console.log('upload file check result: ', pass)

  return pass
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}
</script>
