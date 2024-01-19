<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="被授权飞机:" prop="authz_PlaneId">
          <el-input v-model.number="formData.authz_PlaneId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="被授权航班:" prop="authz_flight">
          <el-input v-model.number="formData.authz_flight" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="权限:" prop="authz_autz">
          <el-input v-model.number="formData.authz_autz" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="授权状态:" prop="authz_state">
          <el-input v-model.number="formData.authz_state" :clearable="false" placeholder="请输入" />
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
  createAuthzPlane,
  updateAuthzPlane,
  findAuthzPlane
} from '@/api/authzPlane'

defineOptions({
    name: 'AuthzPlaneForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            authz_PlaneId: 0,
            authz_flight: 0,
            authz_autz: 0,
            authz_state: 0,
        })
// 验证规则
const rule = reactive({
               authz_PlaneId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authz_flight : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authz_autz : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authz_state : [{
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
      const res = await findAuthzPlane({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reauthzPlane
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAuthzPlane(formData.value)
               break
             case 'update':
               res = await updateAuthzPlane(formData.value)
               break
             default:
               res = await createAuthzPlane(formData.value)
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
