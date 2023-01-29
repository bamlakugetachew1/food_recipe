<script>
import navbar from "./navbar.vue";
import products from "./products.vue";
import { useQuery } from "@vue/apollo-composable";
import { RETRIVE_ALL_RECIPES_QUERY,
        FILTER_BY_TIME,FILTER_BY_TITLE,
        BROWSEBY_TITLE_CREATORSNAME_OR_CATAGORY} from "../graphql/graphql";

export default {
  name: "home",
  components: {
    navbar,
    products,
  },
  data() {
    return {
      data: [],
      pages:1,
      inputdata:""
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
    queryexecute(possiblesortofactions){
      const { result } = useQuery(possiblesortofactions,{
        limit: 6 * this.pages,
       });
       this.data = result;   
     },
    getallrecipes(pages) {
      const { result } = useQuery(RETRIVE_ALL_RECIPES_QUERY,{
        limit: 6 * pages,
      });
      this.data = result;
    },
    isDisabled(e){
     if(e.target.value == "Filter by time"){
      this.queryexecute(FILTER_BY_TIME);
      }
      if(e.target.value == "Filter by title"){
      this.queryexecute(FILTER_BY_TITLE);
      }
      if(e.target.value == "ALL Foods"){
      this.queryexecute(RETRIVE_ALL_RECIPES_QUERY);
      }
      },
      searchrecipes(){
      this.inputdata= "%" + this.inputdata + "%";
      const { result } = useQuery(BROWSEBY_TITLE_CREATORSNAME_OR_CATAGORY,{
        limit: 6 * this.pages,
        search: this.inputdata
        });
        this.data = result;
        this.inputdata = "";
       }
}
};
</script>
<template>
  <div class="w-full">
    <navbar />
    <div class="mb-2 ml-9 flex gap-3" id="shift">
      <select @change="isDisabled">
        <option>ALL Foods</option>
        <option>Filter by time</option>
        <option>Filter by title</option>
      </select>
      <div class="flex gap-1">
      <input v-model="inputdata" class="border rounded overflow-hidden flex px-4 py-2" type="text"  placeholder="Search...">
      <button @click="searchrecipes" class="flex items-center justify-center px-4 border-l">
      <svg class="h-4 w-4 font-bold text-7xl text-grey-dark" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z"/></svg>
      </button>
      </div>
      </div>

     <div v-if="this.data.recipes.length == 0" class="capitalize text-center font-serif font-semibold">there is not matching data try another or refresh page</div>
     <div class="flex gap-1 flex-wrap justify-between">
         <products v-for="product in this.data.recipes" :key="product.id" 
                :id="product.id"
                :title="product.title"
                :imageurl="product.images[0].image_link"
                :fullimagelink="product.images[0].fullimagelink"
            />
   
    </div>
    <a
      href="#"
      @click="loadmore"
      class="mb-3 ml-1 content-center inline-flex items-center px-3 py-2 mt-3 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
    >
      Load More
    </a>
  </div>
</template>
<style scoped></style>
