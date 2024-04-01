<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属用户:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属交易所:" prop="exid">
          <el-input v-model.number="formData.exid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交易对   BTC_USDT:" prop="symbol">
          <el-input v-model="formData.symbol" :clearable="true"  placeholder="请输入交易对   BTC_USDT" />
       </el-form-item>
        <el-form-item label="价格小数位数:" prop="pricePrecision">
          <el-input v-model.number="formData.pricePrecision" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数量小数位数:" prop="volumePrecision">
          <el-input v-model.number="formData.volumePrecision" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ak字段:" prop="ak">
          <el-input v-model="formData.ak" :clearable="true"  placeholder="请输入ak字段" />
       </el-form-item>
        <el-form-item label="sk字段:" prop="sk">
          <el-input v-model="formData.sk" :clearable="true"  placeholder="请输入sk字段" />
       </el-form-item>
        <el-form-item label="对敲:" prop="beat">
          <el-input v-model="formData.beat" :clearable="true"  placeholder="请输入对敲" />
       </el-form-item>
        <el-form-item label="趋势:" prop="trend">
          <el-input v-model="formData.trend" :clearable="true"  placeholder="请输入趋势" />
       </el-form-item>
        <el-form-item label="跟随:" prop="follow">
          <el-input v-model="formData.follow" :clearable="true"  placeholder="请输入跟随" />
       </el-form-item>
        <el-form-item label="盘口自动挂单撤单设置:" prop="pankou">
          <el-input v-model="formData.pankou" :clearable="true"  placeholder="请输入盘口自动挂单撤单设置" />
       </el-form-item>
        <el-form-item label="添加时间:" prop="addtime">
          <el-input v-model.number="formData.addtime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="status字段:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createRobPair,
  updateRobPair,
  findRobPair
} from '@/api/robPair'

defineOptions({
    name: 'RobPairForm'
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
            id: 0,
            uid: 0,
            exid: 0,
            symbol: '',
            pricePrecision: 0,
            volumePrecision: 0,
            ak: '',
            sk: '',
            beat: '',
            trend: '',
            follow: '',
            pankou: '',
            addtime: 0,
            status: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRobPair({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerobPair
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
               res = await createRobPair(formData.value)
               break
             case 'update':
               res = await updateRobPair(formData.value)
               break
             default:
               res = await createRobPair(formData.value)
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
