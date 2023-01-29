<script>
import navbar from './navbar.vue'
import products from './products.vue'
import { useQuery } from "@vue/apollo-composable";
import { RETRIVE_ALL_RECIPES_QUERY } from "../graphql/graphql";

 export default {
  name: 'adminpage',
  components: {
      navbar,
      products
  },
  data() {
    return {
      data: [],
      pages:1
    };
  },
  created() {
    this.getallrecipes(this.pages);
  },
  methods: {
    loadmore(){
        this.pages = this.pages + 1;
        this.getallrecipes(this.pages);
    },
    getallrecipes(pages) {
      const { result } = useQuery(RETRIVE_ALL_RECIPES_QUERY,{
        limit: 6 * pages,
      });
      this.data = result;
     },
    },
  mounted(){
    setInterval( async function () {
      if(localStorage.getItem("itemsdeleted") == "true"){
         localStorage.setItem("itemsdeleted","false");
         location.reload();
      }
    }, 2000);
  },
 }
</script>

<template>
  <div class="w-full">
    <navbar isadminpage="true" />
    <div class="flex gap-1 flex-wrap justify-between">
      <products v-for="product in this.data.recipes" :key="product.id" 
                :id="product.id"
                :title="product.title"
                :imageurl="product.images[0].image_link"
                :fullimagelink="product.images[0].fullimagelink"
                isadminpage="true" 
      />
    </div>
    <a       
     @click="loadmore"
     href="#" class=" mb-3 ml-1 content-center inline-flex items-center px-3 py-2 mt-3 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
     Load More
   </a>
  </div>
</template>
<style scoped>

</style>
