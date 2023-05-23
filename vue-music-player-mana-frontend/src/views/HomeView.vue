<template>
  <div>
    <p>this is homeview</p>
    <button @click="getAsync">异步</button>
    <button @click="getsync">同步</button>
    <p>{{ str }}</p>
  </div>
</template>

<script setup lang="ts">
import axios from 'axios'
import { ref } from 'vue'

const str = ref({
  Username :"",
	Gender  :"",
	Age      :"",
	Email      :"",
	Password   :"",
})
//异步回调
function getAsync() {
  axios.get('http://localhost:8080/hello').then(function (response: any) {
    //回调
    // str.value.name = response
  })
}

//同步方式
async function getsync() {
  await axios.post('http://localhost:8080/register',{
  Username :"asdasd",
	Gender  :"asdasd",
	Age      :"ssssss",
	Email      :"asdasdsa",
	Password   :"asdsdsdsd",
}).then(function (response: any) {
    //回调
    try {
      str.value = response
    } catch (error: any) {
      console.log(error)
    }
  })
}
</script>
