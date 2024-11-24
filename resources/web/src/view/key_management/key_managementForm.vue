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
          label="ID:"
          prop="key_id"
        >
          <el-input
            v-model="formData.key_id"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="密钥类型:"
          prop="kind"
        >
          <el-input
            v-model="formData.kind"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="所有者1:"
          prop="user1"
        >
          <el-input
            v-model.number="formData.user1"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="所有者2:"
          prop="user2"
        >
          <el-input
            v-model.number="formData.user2"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="密钥长度:"
          prop="length"
        >
          <el-input
            v-model.number="formData.length"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="密钥状态:"
          prop="key_status"
        >
          <el-select
            v-model="formData.key_status"
            placeholder="请选择"
            style="width:100%"
            :clearable="true"
          >
            <el-option
              v-for="item in []"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="更新间隔:"
          prop="update_time"
        >
          <el-input
            v-model.number="formData.update_time"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="密钥密文:"
          prop="ciphertext"
        >
          <el-input
            v-model="formData.ciphertext"
            :clearable="true"
            placeholder="请输入"
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
  createKeyEntity,
  updateKeyEntity,
  findKeyEntity
} from '@/api/key_management'

defineOptions({
  name: 'KeyEntityForm'
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
  key_id: '',
  kind: '',
  user1: 0,
  user2: 0,
  length: 0,
  update_time: 0,
  ciphertext: '',
})
// 验证规则
const rule = reactive({
  key_id: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  kind: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  user1: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  user2: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  length: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  key_status: [{
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
    const res = await findKeyEntity({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rekm
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
        res = await createKeyEntity(formData.value)
        break
      case 'update':
        res = await updateKeyEntity(formData.value)
        break
      default:
        res = await createKeyEntity(formData.value)
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
