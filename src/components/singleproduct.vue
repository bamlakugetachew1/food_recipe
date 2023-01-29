<script>
 import { ref } from "vue";
 import { useQuery, useMutation} from "@vue/apollo-composable";
 import { RETRIVE_SINGLE_RECIPES_QUERY,
          RETRIVE_COMMENTS,CREATE_COMMENTS,
          INSERT_RATINGS,
          RETRIVE_AGGREGATE_RATING_RECIPES_QUERY
        } from "../graphql/graphql";
 import {checkuserloggedin} from "../helper/helper";
 export default {
  name: 'singleproduct',
  components: {
  },
  created(){
     this.id = this.$route.params.id;
     this.userid = localStorage.getItem("userid");
     this.getonerecipes(this.id);
     this.getonerecipecomments(this.id);
     this.avgrating(this.id);
  },
  data(){
    return{
        id:0,
        data:{},
        comments:{},
        message:"",
        userid:"",
        aggregateval:{},
    }
  },
  methods:{
    getonerecipes(id) {
      const { result } = useQuery(RETRIVE_SINGLE_RECIPES_QUERY,{
        id: id,
      });
      this.data = result;
    },

     getonerecipecomments(recipesid) {
      const { result } = useQuery(RETRIVE_COMMENTS,{
        recipes_id: recipesid
      });
      this.comments = result;
    },

     review(){
        checkuserloggedin();
        this.addreview({ comment: this.message, user: this.userid,recipeid:this.id }).then(
        (data) => {
            console.log(data);
        }
      );
       this.message = "";
       this.getonerecipecomments(this.id);
     },

      rating(rate){
        checkuserloggedin();
        this.addrating({ rate: rate, user: this.userid,recipeid:this.id }).then(
        (data) => {
          this.addingstar(rate);     
          }
         );  
       },

     addingstar(ratevalue){
        if(ratevalue >= 1){
        let ratecount = Math.floor(ratevalue);
         for(let i=1; i<=5; i++){
         document.getElementById(i).classList.remove("text-yellow-400");  
         }
         for(let i=1; i<=ratecount; i++){
         document.getElementById(i).classList.add("text-yellow-400");
         }  
        } 
     },

     avgrating(recipeid){
      const { result } = useQuery(RETRIVE_AGGREGATE_RATING_RECIPES_QUERY,{
        recipeid: recipeid,
      });
      this.aggregateval = result; 
      }
  },
  setup() {
    const comment = ref("");
    const user = ref(0);
    const recipeid = ref(0);
    const rate = ref(0);

    const { mutate: addreview } = useMutation(CREATE_COMMENTS,
      () => ({
        variables: {
         comment: comment.value,
         user: user.value,
         recipeid: recipeid.value,

        },
      })
    );
    
    const { mutate: addrating } = useMutation(INSERT_RATINGS,
      () => ({
        variables: {
         rate: rate.value,
         user: user.value,
         recipeid: recipeid.value,

        },
      })
    );

    return {
       addreview,
       addrating
    };
  },
 }
</script>
<template>
 <div class="md:flex items-start justify-center py-12 2xl:px-20 md:px-6 px-4">
            <div class="xl:w-2/6 lg:w-2/5 w-80 md:block hidden">
                <div v-for="imagelink in this.data.recipes_by_pk.images" :key="imagelink.id">
            
                <img class="w-full mt-2" alt="image of a girl posing" v-bind:src="'http://localhost:8000/images/' +  imagelink.image_link" />
                
                </div>
                  </div> 
            <div class="md:hidden">
                <img class="w-full" alt="image of a girl posing" v-bind:src="'http://localhost:8000/images/' +  this.data.recipes_by_pk.images[0].image_link" />
                <div class="flex items-center justify-between mt-3 space-x-4 md:space-x-0">
                    <div v-for="imagelink in this.data.recipes_by_pk.images" :key="imagelink.id">
                    <img alt="image-tag-one" class="md:w-48 md:h-48 w-full" v-bind:src="'http://localhost:8000/images/' +  imagelink.image_link" />

                    </div>
                    </div>
            </div>
            <div class="xl:w-2/5 md:w-1/2 lg:ml-8 md:ml-6 md:mt-0 mt-6">
                <div class="border-b border-gray-200 pb-6">
                    <h1 class="leading-none text-xl font-semibold text-gray-600 dark:text-gray-300 "> {{this.data.recipes_by_pk.title}}</h1>
                    <h1 class="lg:text-2xl text-xl font-semibold lg:leading-6 leading-7 text-gray-800 dark:text-white mt-2">{{this.data.recipes_by_pk.catagories}}</h1>
                </div>
             <div class="flex items-center mt-2">
             <svg @click="rating(1)" id="1" aria-hidden="true" class="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>First star</title><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
             <svg @click="rating(2)" id="2" aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Second star</title><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
             <svg @click="rating(3)" id="3" aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Third star</title><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
             <svg @click="rating(4)" id="4" aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Fourth star</title><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
             <svg @click="rating(5)" id="5" aria-hidden="true" class="w-5 h-5 dark:text-gray-500" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Fifth star</title><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
             <p class="ml-2 text-sm font-medium text-gray-500 dark:text-gray-400">
                {{ this.aggregateval.count_ratings_aggregate.aggregate.avg.rate}}
                 out of 5 in  
                {{ this.aggregateval.count_ratings_aggregate.aggregate.count}}
                 Ratings
                </p>
             </div>
                <div>
                    <p class="xl:pr-48 text-base lg:leading-tight leading-normal text-gray-600 dark:text-gray-300 mt-7">{{this.data.recipes_by_pk.description}}</p>
                    <p class="text-base leading-4 mt-7 text-gray-600 dark:text-gray-300">Time: {{this.data.recipes_by_pk.time}}min</p>
                    <p class="text-base leading-4 mt-4 text-gray-600 dark:text-gray-300">Creators name : {{this.data.recipes_by_pk.creatorsname}}</p>
                   </div>
                <div>
                    <div class="border-t border-b py-4 mt-7 border-gray-200">
                        <div data-menu class="flex justify-between items-center cursor-pointer">
                            <p class="text-base leading-4 text-gray-800 dark:text-gray-300">Ingridents</p> 
                        </div>
                    </div>

                        <div v-for="ingridents in this.data.recipes_by_pk.ingridents"  class="mt-2 flex  justify-between">
                          <p class="font-serif font-semibold">{{ ingridents.ingrident}}</p>
                        </div>

                    <div class="border-t border-b py-4 mt-7 border-gray-200">
                        <div data-menu class="flex justify-between items-center cursor-pointer">
                            <p class="text-base leading-4 text-gray-800 dark:text-gray-300">Steps To prepare</p> 
                        </div>
                    </div>
                    <div class="mt-2 flex flex-wrap justify-between" v-for="steps in this.data.recipes_by_pk.steps"  >
                          <p class="font-serif font-semibold">{{ steps.step}}</p>
                        </div>

                    <div class="border-t border-b py-4 mt-7 border-gray-200">
                        <div data-menu class="flex justify-between items-center cursor-pointer">
                            <p class="text-base leading-4 text-gray-800 dark:text-gray-300">Reviews and Ratings</p> 
                        </div>
                    </div>
                </div>
                <div>

                    <div class="border-b py-4 border-gray-200" v-for="comments in this.comments.comment">
                        <p class="md:w-96 text-base leading-normal text-gray-600 dark:text-gray-300 mt-4">{{comments.comment}}</p>
                        </div>
                    <div>
                    <div class="border-t border-b py-4 mt-7 border-gray-200">
                        <div data-menu class="flex justify-between items-center cursor-pointer">
                            <p class="text-base leading-4 text-gray-800 dark:text-gray-300">Write Your Review</p> 
                        </div>
                    </div>
                    <div class="flex mt-2">
                        <input v-model="message" type="text" id="large-input" class="block w-full p-4 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-md focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                        <button @click="review" type="button" class="text-white ml-2 bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Review</button>
                    </div>
                    
                </div>            
            </div>
            </div>
        </div>  
  </template>
  <style>
  </style>