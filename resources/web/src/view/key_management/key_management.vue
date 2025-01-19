<template>
  <div>
    <!-- <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="ID"
          prop="id"
        >
          <el-input
            v-model="searchInfo.id"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="密钥类型"
          prop="kind"
        >
          <el-input
            v-model="searchInfo.kind"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="所有者1"
          prop="user1"
        >

          <el-input
            v-model.number="searchInfo.user1"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="所有者2"
          prop="user2"
        >

          <el-input
            v-model.number="searchInfo.user2"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="密钥状态"
          prop="key_status"
        >
          <el-input
            v-model="searchInfo.key_status"
            placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
          label="密钥密文"
          prop="ciphertext"
        >
          <el-input
            v-model="searchInfo.ciphertext"
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
    </div> -->
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
          label="密钥ID"
          width="180"
          prop="id"
        >
        </el-table-column>
        <el-table-column
          align="left"
          label="密钥类型"
          prop="key_type"
          width="120"
        />
        <el-table-column
          align="left"
          label="所有者1"
          prop="owner1"
          width="120"
        />
        <el-table-column
          align="left"
          label="所有者2"
          prop="owner2"
          width="120"
        />
        <el-table-column
          align="left"
          label="密钥长度"
          prop="key_len"
          width="120"
        />
        <el-table-column
          align="left"
          label="密钥状态"
          prop="key_state"
          width="120"
        />
    
        <el-table-column
          align="left"
          label="更新间隔"
          prop="updatecycle"
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
              @click="updateKeyEntityFunc(scope.row)"
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
            label="密钥类型:"
            prop="kind"
          >
            <el-select
              v-model="formData.key_type"
              filterable
              placeholder="请选择密钥类型"
              style="width:100%"
              :clearable="true"
            >
              <el-option
                v-for="(item, key) in keyTypeOptions"
                :key="key"
                :label="item.label"
                :value="item.label"
              />
            </el-select>

          </el-form-item>
          <el-form-item
            label="AS UA:"
            prop="user1"
          >
            <el-select
              v-model="formData.owner1"
              filterable
              placeholder="请选择"
              style="width:100%"
              :clearable="true"
            >
              <el-option
                v-for="(item,key) in plane_opts"
                :key="key"
                :label="item.company + ' - ' + item.plane_id + ' - ' +item.ua"
                :value="item.ua.toString()"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="SGW ID:"
            prop="user2"
          >
            <el-input
              v-model="formData.owner2"
              :clearable="true"
              placeholder="请输入所有者2"
            />
          </el-form-item>
          <el-form-item
            label="密钥长度:"
            prop="length"
          >
            <el-input
              v-model.number="formData.key_len"
              :clearable="true"
              placeholder="请输入密钥长度"
            />
          </el-form-item>
          <el-form-item
            label="更新间隔:"
            prop="update_time"
          >
            <el-input
              v-model.number="formData.updatecycle"
              :clearable="true"
              placeholder="请输入更新间隔"
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
          <el-descriptions-item label="ID">
            {{ formData.key_id }}
          </el-descriptions-item>
          <el-descriptions-item label="密钥类型">
            {{ formData.kind }}
          </el-descriptions-item>
          <el-descriptions-item label="所有者1">
            {{ formData.user1 }}
          </el-descriptions-item>
          <el-descriptions-item label="所有者2">
            {{ formData.user2 }}
          </el-descriptions-item>
          <el-descriptions-item label="密钥长度">
            {{ formData.length }}
          </el-descriptions-item>
          <el-descriptions-item label="密钥状态">
            {{ formData.key_status }}
          </el-descriptions-item>
          <el-descriptions-item label="更新间隔">
            {{ formData.update_time }}
          </el-descriptions-item>
          <el-descriptions-item label="密钥密文">
            {{ formData.ciphertext }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>
<script>
  export default {
    name: 'Key_management',
    data(){
    return{
            options:[{
                value:1,
                label:1
            },
            {
                value:7,
                label:7 
            },
            {
                value:30,
                label:30 
            },
            {
                value:180,
                label:180 
            },
            {
                value:365,
                label:365
            }],
            options5:[{
                value:'AS1',
                label:'AS1'
            },{
                value:'AS2',
                label:'AS2'
            },{
                value:'Berry',
                label:'Berry'
            }],
            options3: [{
               value: 16,
               label: 16
             }, {
              value: 32,
              label: 32
            }],
           options4: [{
            value: 'ROOT_KEY',
            label: 'ROOT_KEY'
            }],
        options6: [{
               value: 'ACTIVE',
               label: 'ACTIVE'
             },{
               value: 'PRE_ACTIVATION',
               label: 'PRE_ACTIVATION'
             }],
            
             options7:[{
                value:'GS1',
                label:'GS1'
            },
            {
                value:'GSt',
                label:'GSt' 
            },{
                value:'SGW',
                label:'SGW' 
            }],
            value: ''
        }
    },
};

</script>
<script setup>
import {
  createKeyEntity,
  deleteKeyEntity,
  deleteKeyEntityByIds,
  updateKeyEntity,
  findKeyEntity,
  getKeyEntityList,
  getOptions
} from '@/api/key_management'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'KeyEntity'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  key_type: '',
  owner1: '',
  owner2: '10000',
  key_len: 16,
  updatecycle: 365,
})
const keyTypeOptions = ref([])
const keyStateOptions = ref([])
const plane_opts = ref([])
// 验证规则
const rule = reactive({
  // key_id: [{
  //   required: true,
  //   message: '',
  //   trigger: ['input', 'blur'],
  // },
  // {
  //   whitespace: true,
  //   message: '不能只输入空格',
  //   trigger: ['input', 'blur'],
  // }
  // ],
  // kind: [{
  //   required: true,
  //   message: '',
  //   trigger: ['input', 'blur'],
  // },
  // {
  //   whitespace: true,
  //   message: '不能只输入空格',
  //   trigger: ['input', 'blur'],
  // }
  // ],
  // user1: [{
  //   required: true,
  //   message: '',
  //   trigger: ['input', 'blur'],
  // },
  // ],
  // user2: [{
  //   required: true,
  //   message: '',
  //   trigger: ['input', 'blur'],
  // },
  // ],
  // length: [{
  //   required: true,
  //   message: '',
  //   trigger: ['input', 'blur'],
  // },
  // ],
  // key_status: [{
  //   required: true,
  //   message: '',
  //   trigger: ['input', 'blur'],
  // },
  // ],
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
  const table = await getKeyEntityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    console.log(tableData.value, total.value)
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  const res = await getOptions()

  if (res.code === 0) {
    // options.value = res.data.options
    plane_opts.value = res.data.options.plane_ids
    console.log(plane_opts.value)
  }
  keyTypeOptions.value = await getDictFunc('KeyType')
  keyStateOptions.value = await getDictFunc('KeyState')
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
    deleteKeyEntityFunc(row)
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
          ids.push(item.id)
        })
  const res = await deleteKeyEntityByIds({ ids })
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
const updateKeyEntityFunc = async(row) => {
  const res = await findKeyEntity({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rekm
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteKeyEntityFunc = async(row) => {
  const res = await deleteKeyEntity({ id: row.id })
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
  const res = await findKeyEntity({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rekm
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    key_id: '',
    kind: '',
    user1: 0,
    user2: 0,
    length: 0,
    update_time: 0,
    ciphertext: '',
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
    key_id: '',
    kind: '',
    user1: 0,
    user2: 0,
    length: 0,
    update_time: 0,
    ciphertext: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    console.log(formData.value)
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
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
