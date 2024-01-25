<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="认证飞机站:" prop="authc_as_sac">
          <el-input v-model.number="formData.authc_as_sac" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前地面站:" prop="authc_gs_sac">
          <el-input v-model.number="formData.authc_gs_sac" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前地面控制站:" prop="authc_gsc_sac">
          <el-input v-model.number="formData.authc_gsc_sac" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="认证状态:" prop="authc_state">
          <el-select v-model="formData.authc_state" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in AuthenticationOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="状态转换时间:" prop="authc_trans_time">
          <el-date-picker v-model="formData.authc_trans_time" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAuthcState,
  updateAuthcState,
  findAuthcState
} from '@/api/authcState'

defineOptions({
    name: 'AuthcStateForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const AuthenticationOptions = ref([])
const formData = ref({
            authc_as_sac: 0,
            authc_gs_sac: 0,
            authc_gsc_sac: 0,
            authc_state: undefined,
            authc_trans_time: new Date(),
        })
// 验证规则
const rule = reactive({
               authc_as_sac : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authc_gs_sac : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authc_gsc_sac : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authc_state : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authc_trans_time : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAuthcState({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reauthcState
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    AuthenticationOptions.value = await getDictFunc('Authentication')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAuthcState(formData.value)
               break
             case 'update':
               res = await updateAuthcState(formData.value)
               break
             default:
               res = await createAuthcState(formData.value)
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
