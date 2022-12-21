<script setup>
import { ref } from "vue"
import ConnectionList from './components/ConnectionList.vue'
import ConnectionManage from './components/ConnectionManage.vue'
import Keys from './components/Keys.vue'
import KeyValue from './components/KeyValue.vue'

let flushFlag = ref(true)
let keyDB = ref()
let keyIdentity = ref()
let keyKey = ref()

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

function selectDB(db, identity) {
  keyDB.value = db
  keyIdentity = identity
}

function selectKey(key) {
  keyKey.value = key
}

</script>

<template>
  <el-row :gutter="0" style="display: flex;">
    <el-col :span="5" style="height: 100vh; padding: 12px;">
      <div style="margin-bottom: 12px;">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList :flush="flushFlag" @emit-select-db="selectDB"/>
    </el-col>
    <el-col :span="6" style="padding: 12px;">
      <Keys :keyDB="keyDB" :keyIdentity="keyIdentity" @emit-select-key="selectKey"/>
    </el-col>
    <el-col :span="12" style="padding: 12px;">
      <KeyValue :keyDB="keyDB" :keyIdentity="keyIdentity" :keyKey="keyKey"/>
    </el-col>
  </el-row>
</template>

<style>
</style>
