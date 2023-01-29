<script>
import { ref } from "vue";
import {useQuery,useMutation} from "@vue/apollo-composable";
import { DELETE_RECIPES,CHECK_ALERADY_LIKED,ADD_USER_LIKES} from "../graphql/graphql";
import {checkuserloggedin} from "../helper/helper";

 export default {
  name: 'products',
  components: { },
  props: {
    isadminpage: Boolean,
    imageurl: String,
    id: String,
    title:String,
    fullimagelink:String
    },
    created(){
       this.userid = localStorage.getItem("userid");
    },
    data(){
      return{
        path:"",
        userid:0,
      }
    },
    methods:{
        deletefoods(id){
          this.deleterecipes({id:id}).then((data)=>{ 
          localStorage.setItem("itemsdeleted","true");
            });
           },
        checkaleradyliked (userid,recipeid) {
          const { result } =  useQuery(CHECK_ALERADY_LIKED,{
          users_id:userid,
          recipes_id:recipeid,
           });
           return   result;
           },
            likefoods(userid,recipeid){
            checkuserloggedin();
            var result = this.checkaleradyliked(userid,recipeid);
                var object = {
                  "userid": userid,
                  "recipeid": recipeid
                }
            if(result){  this.addlikes({objects:object}).then((data)=>{ console.log(data) });}
             }
             },

    setup() {
    const   id = ref(0);
    const   object = ref({});
    const { mutate: deleterecipes } = useMutation(DELETE_RECIPES,() => ({
      variables: {
        id: id.value
      },
    }));
     const { mutate: addlikes } = useMutation(ADD_USER_LIKES,() => ({
      variables: {
        objects:object.value
      },
    }));
  
    return {
      deleterecipes,
      addlikes
    };
  }
 }
</script>
<template>
<div class="max-w-sm  bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
    <a href="#">
        <img class="rounded-t-lg" v-bind:src="fullimagelink" alt="" />
    </a>
    <div class="p-5">
       <div class="flex justify-between">
        <p class="font-serif ">{{title}}</p>
      </div>
      <div class="flex justify-between mt-1">
        <button @click="likefoods(this.userid,id)" v-if="!isadminpage"> <i  class="fa fa-heart"  aria-hidden="true" id="space" ></i></button>
        <button v-if="isadminpage">  <i  class="fa fa-edit"  aria-hidden="true" id="space" ></i></button>
        <button @click="deletefoods(id)" v-if="isadminpage"> <i  class="fa fa-trash"  aria-hidden="true" id="space" ></i></button>
         
         <router-link v-if="!isadminpage"
          :to="{  path:`/singleproduct/${id}`, params: {productid:id}}"
          >
          <a  href="#" class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            More
         <svg aria-hidden="true" class="w-4 h-4 ml-2 -mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
          </a>
          </router-link>
          </div>
     
    </div>
</div>
</template>
<style scoped>

</style>