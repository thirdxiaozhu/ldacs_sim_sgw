<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item
          label="地面站SAC:"
          prop="gs_sac"
        >
          <el-input
            v-model.number="formData.gs_sac"
            :clearable="false"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="北纬:"
          prop="latitude_n"
        >
          <el-input-number
            v-model="formData.latitude_n"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item
          label="东经:"
          prop="longtitude_e"
        >
          <el-input-number
            v-model="formData.longtitude_e"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="save"
          >保存</el-button>
          <el-button
            type="primary"
            @click="back"
          >返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAccountGs,
  updateAccountGs,
  findAccountGs
} from '@/api/accountGs'

defineOptions({
  name: 'AccountGsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  gs_sac: 0,
  latitude_n: 0,
  longtitude_e: 0,
})
// 验证规则
const rule = reactive({
  gs_sac: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findAccountGs({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reaccountGs
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createAccountGs(formData.value)
        break
      case 'update':
        res = await updateAccountGs(formData.value)
        break
      default:
        res = await createAccountGs(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
