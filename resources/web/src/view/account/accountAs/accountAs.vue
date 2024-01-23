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
          label="飞机注册号"
          prop="as_plane_id"
        >

          <el-input
            v-model.number="searchInfo.as_plane_id"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="执飞航班号"
          prop="as_flight"
        >

          <el-input
            v-model.number="searchInfo.as_flight"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="执飞日期"
          prop="as_date"
        >

          <template #label>
            <span>
              执飞日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startAsDate"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="time=> searchInfo.endAsDate ? time.getTime() > searchInfo.endAsDate.getTime() : false"
          />
          —
          <el-date-picker
            v-model="searchInfo.endAsDate"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="time=> searchInfo.startAsDate ? time.getTime() < searchInfo.startAsDate.getTime() : false"
          />

        </el-form-item>
        <el-form-item
          label="飞机站SAC"
          prop="as_sac"
        >
          <el-input
            v-model="searchInfo.as_sac"
            placeholder="搜索条件"
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
          label="飞机注册号"
          prop="plane_id.plane_id"
          width="120"
        />
        <el-table-column
          align="left"
          label="执飞航班号"
          prop="flight.flight"
          width="120"
        />
        <el-table-column
          align="left"
          label="执飞日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.as_date) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="飞机站SAC"
          prop="as_sac"
          width="120"
        />
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
              @click="updateAccountAsFunc(scope.row)"
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
            label="飞机注册号:"
            prop="as_plane_id"
          >
            <el-select
              v-model="formData.as_plane_id"
              filterable
              placeholder="请选择"
              style="width:100%"
              :clearable="true"
            >
              <el-option
                v-for="(item, key) in plane_opts"
                :key="key"
                :label="item.plane_id"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="执飞航班号:"
            prop="as_flight"
          >
            <el-select
              v-model="formData.as_flight"
              filterable
              placeholder="请选择"
              style="width:100%"
              :clearable="true"
            >
              <el-option
                v-for="(item,key) in flight_opts"
                :key="key"
                :label="item.flight"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="执飞日期:"
            prop="as_date"
          >
            <el-date-picker
              v-model="formData.as_date"
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
          <el-descriptions-item label="飞机注册号">
            {{ formData.as_plane_id }}
          </el-descriptions-item>
          <el-descriptions-item label="执飞航班号">
            {{ formData.as_flight }}
          </el-descriptions-item>
          <el-descriptions-item label="执飞日期">
            {{ formatDate(formData.as_date) }}
          </el-descriptions-item>
          <el-descriptions-item label="飞机站SAC">
            {{ formData.as_sac }}
          </el-descriptions-item>
          <el-descriptions-item label="当前GS">
            {{ formData.state.gs_sac }}
          </el-descriptions-item>
          <el-descriptions-item label="当前GSC">
            {{ formData.state.gsc_sac }}
          </el-descriptions-item>
          <el-descriptions-item label="认证套件">
            {{ formData.state.auth_id }}
          </el-descriptions-item>
          <el-descriptions-item label="认证状态">
            {{ formData.state.auth_state }}
          </el-descriptions-item>
          <el-descriptions-item label="加密套件">
            {{ formData.state.enc_id }}
          </el-descriptions-item>
          <el-descriptions-item label="是否认证成功">
            {{ formData.state.is_success }}
          </el-descriptions-item>
          <el-descriptions-item label="是否结束">
            {{ formData.state.is_term }}
          </el-descriptions-item>
          <el-descriptions-item label="KDF">
            {{ formData.state.kdf_k }}
          </el-descriptions-item>
          <el-descriptions-item label="随机数">
            {{ formData.state.rand_v }}
          </el-descriptions-item>
          <el-descriptions-item label="共享密钥">
            {{ formData.state.shared_key }}
          </el-descriptions-item>
          <el-descriptions-item label="SNP状态">
            {{ formData.state.snp_state }}
          </el-descriptions-item>
          <el-descriptions-item label="序列号">
            {{ formData.state.sqn }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createAccountAs,
  deleteAccountAs,
  deleteAccountAsByIds,
  updateAccountAs,
  findAccountAs,
  getAccountAsList,
  getOptions,
  setStateChange,
} from '@/api/accountAs'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'AccountAs'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  as_plane_id: 0,
  as_flight: 0,
  as_date: new Date(),
  as_sac: 0,
  state: '',
})

// 验证规则
const rule = reactive({
  as_plane_id: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  as_flight: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  as_date: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  as_sac: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
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
  as_date: [{ validator: (rule, value, callback) => {
    if (searchInfo.value.startAsDate && !searchInfo.value.endAsDate) {
      callback(new Error('请填写结束日期'))
    } else if (!searchInfo.value.startAsDate && searchInfo.value.endAsDate) {
      callback(new Error('请填写开始日期'))
    } else if (searchInfo.value.startAsDate && searchInfo.value.endAsDate && (searchInfo.value.startAsDate.getTime() === searchInfo.value.endAsDate.getTime() || searchInfo.value.startAsDate.getTime() > searchInfo.value.endAsDate.getTime())) {
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
  const table = await getAccountAsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

const plane_opts = ref([])
const flight_opts = ref([])
const auth_opts = ref([])
// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  const res = await getOptions()

  console.log(res)
  if (res.code === 0) {
    // options.value = res.data.options
    plane_opts.value = res.data.options.plane_ids
    flight_opts.value = res.data.options.flights
    auth_opts.value = res.data.options.authzs
  }
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
    deleteAccountAsFunc(row)
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
  const res = await deleteAccountAsByIds({ ids })
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
const updateAccountAsFunc = async(row) => {
  const res = await findAccountAs({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reaccountAs
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAccountAsFunc = async(row) => {
  const res = await deleteAccountAs({ ID: row.ID })
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
  const res = await findAccountAs({ ID: row.ID })
  if (res.code === 0) {
    console.log(res.data)
    formData.value = res.data.reaccountAs
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    as_plane_id: 0,
    as_flight: 0,
    as_date: new Date(),
    state: '',
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
    as_plane_id: 0,
    as_flight: 0,
    as_date: new Date(),
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createAccountAs(formData.value)
        break
      case 'update':
        res = await updateAccountAs(formData.value)
        break
      default:
        res = await createAccountAs(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      await getTableData()
    }
  })
}

</script>

<style>

</style>
