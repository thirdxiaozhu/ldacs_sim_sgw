<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="飞机注册号:" prop="as_plane_id">
          <el-input v-model.number="formData.as_plane_id" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="航班号:" prop="as_flight">
          <el-input v-model.number="formData.as_flight" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执飞日期:" prop="flight_date">
          <el-date-picker v-model="formData.flight_date" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
       </el-form-item>
        <el-form-item label="飞机站SAC:" prop="as_sac">
          <el-input v-model="formData.as_sac" :clearable="false" placeholder="请输入" />
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
  createAirStation,
  updateAirStation,
  findAirStation
} from '@/api/airStation'

defineOptions({
  name: 'AirStationForm'
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
            as_plane_id: 0,
            as_flight: 0,
            flight_date: new Date(),
            as_sac: '',
        })
// 验证规则
const rule = reactive({
               as_plane_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               as_flight : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               flight_date : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               as_sac : [{
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
      const res = await findAirStation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reairStation
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
               res = await createAirStation(formData.value)
               break
             case 'update':
               res = await updateAirStation(formData.value)
               break
             default:
               res = await createAirStation(formData.value)
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
