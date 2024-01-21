<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="创建日期"
          prop="createdAt"
        >
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"
          />
        </el-form-item>
        <el-form-item
          label="认证飞机站"
          prop="authc_as_sac"
        >

          <el-input
            v-model.number="searchInfo.authc_as_sac"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="当前地面站"
          prop="authc_gs_sac"
        >

          <el-input
            v-model.number="searchInfo.authc_gs_sac"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="当前地面控制站"
          prop="authc_gsc_sac"
        >

          <el-input
            v-model.number="searchInfo.authc_gsc_sac"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="认证状态"
          prop="authc_state"
        >
          <el-select
            v-model="searchInfo.authc_state"
            clearable
            placeholder="请选择"
            @clear="()=>{searchInfo.authc_state=undefined}"
          >
            <el-option
              v-for="(item,key) in AuthenticationOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="状态转换时间"
          prop="authc_trans_time"
        >

          <template #label>
            <span>
              状态转换时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startAuthcTransTime"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="time=> searchInfo.endAuthcTransTime ? time.getTime() > searchInfo.endAuthcTransTime.getTime() : false"
          />
          —
          <el-date-picker
            v-model="searchInfo.endAuthcTransTime"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="time=> searchInfo.startAuthcTransTime ? time.getTime() < searchInfo.startAuthcTransTime.getTime() : false"
          />

        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
              type="primary"
              link
              @click="deleteVisible = false"
            >取消</el-button>
            <el-button
              type="primary"
              @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="认证飞机站"
          prop="authc_as_sac"
          width="120"
        />
        <el-table-column
          align="left"
          label="当前地面站"
          prop="authc_gs_sac"
          width="120"
        />
        <el-table-column
          align="left"
          label="当前地面控制站"
          prop="authc_gsc_sac"
          width="120"
        />
        <el-table-column
          align="left"
          label="认证状态"
          prop="authc_state"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.authc_state,AuthenticationOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="状态转换时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.authc_trans_time) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          min-width="120"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateAuthcStateFunc(scope.row)"
            >变更</el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type==='create'?'添加':'修改'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="80px"
        >
          <el-form-item
            label="认证飞机站:"
            prop="authc_as_sac"
          >
            <el-input
              v-model.number="formData.authc_as_sac"
              :clearable="false"
              placeholder="请输入认证飞机站"
            />
          </el-form-item>
          <el-form-item
            label="当前地面站:"
            prop="authc_gs_sac"
          >
            <el-input
              v-model.number="formData.authc_gs_sac"
              :clearable="false"
              placeholder="请输入当前地面站"
            />
          </el-form-item>
          <el-form-item
            label="当前地面控制站:"
            prop="authc_gsc_sac"
          >
            <el-input
              v-model.number="formData.authc_gsc_sac"
              :clearable="false"
              placeholder="请输入当前地面控制站"
            />
          </el-form-item>
          <el-form-item
            label="认证状态:"
            prop="authc_state"
          >
            <el-select
              v-model="formData.authc_state"
              placeholder="请选择认证状态"
              style="width:100%"
              :clearable="false"
            >
              <el-option
                v-for="(item,key) in AuthenticationOptions"
                :key="key"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="状态转换时间:"
            prop="authc_trans_time"
          >
            <el-date-picker
              v-model="formData.authc_trans_time"
              type="date"
              style="width:100%"
              placeholder="选择日期"
              :clearable="false"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="认证飞机站">
            {{ formData.authc_as_sac }}
          </el-descriptions-item>
          <el-descriptions-item label="当前地面站">
            {{ formData.authc_gs_sac }}
          </el-descriptions-item>
          <el-descriptions-item label="当前地面控制站">
            {{ formData.authc_gsc_sac }}
          </el-descriptions-item>
          <el-descriptions-item label="认证状态">
            {{ filterDict(formData.authc_state,AuthenticationOptions) }}
          </el-descriptions-item>
          <el-descriptions-item label="状态转换时间">
            {{ formatDate(formData.authc_trans_time) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createAuthcState,
  deleteAuthcState,
  deleteAuthcStateByIds,
  updateAuthcState,
  findAuthcState,
  getAuthcStateList
} from '@/api/authcState'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'AuthcState'
})

// 自动化生成的字典（可能为空）以及字段
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
  authc_as_sac: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  authc_gs_sac: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  authc_gsc_sac: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  authc_state: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  authc_trans_time: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
  authc_trans_time: [{ validator: (rule, value, callback) => {
    if (searchInfo.value.startAuthcTransTime && !searchInfo.value.endAuthcTransTime) {
      callback(new Error('请填写结束日期'))
    } else if (!searchInfo.value.startAuthcTransTime && searchInfo.value.endAuthcTransTime) {
      callback(new Error('请填写开始日期'))
    } else if (searchInfo.value.startAuthcTransTime && searchInfo.value.endAuthcTransTime && (searchInfo.value.startAuthcTransTime.getTime() === searchInfo.value.endAuthcTransTime.getTime() || searchInfo.value.startAuthcTransTime.getTime() > searchInfo.value.endAuthcTransTime.getTime())) {
      callback(new Error('开始日期应当早于结束日期'))
    } else {
      callback()
    }
  }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAuthcStateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  console.log(table.data)
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  AuthenticationOptions.value = await getDictFunc('Authentication')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteAuthcStateFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteAuthcStateByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAuthcStateFunc = async(row) => {
  const res = await findAuthcState({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reauthcState
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAuthcStateFunc = async(row) => {
  const res = await deleteAuthcState({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findAuthcState({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reauthcState
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    authc_as_sac: 0,
    authc_gs_sac: 0,
    authc_gsc_sac: 0,
    authc_state: undefined,
    authc_trans_time: new Date(),
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    authc_as_sac: 0,
    authc_gs_sac: 0,
    authc_gsc_sac: 0,
    authc_state: undefined,
    authc_trans_time: new Date(),
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
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
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
