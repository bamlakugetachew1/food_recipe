<script>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import { INSERT_RECIPES_ONE,
    INSERT_RECIPES_INGRIDENTS,
    INSERT_RECIPES_STEPS,
    INSERT_RECIPES_IMAGES} from "../graphql/graphql";
import router from "../router/index";
import axios from "axios";
export default {
  name: "addproduct",
  components: {},
  created() {
    this.creatorsname = localStorage.getItem("username");
    this.creatorsid =   localStorage.getItem("userid");
  },
  data() {
    return {
      title: "",
      description: "",
      catagories: "",
      time: "",
      creatorsname: "",
      creatorsid: "",
      ingridents: "",
      steps: "",
      selectedFile: null,
      filelocations:""

    };
  },
  methods: {
    onFileChange(e) {
      this.selectedFile = e.target.files;     
    },
   async  onUploadFile() {
      const formData = new FormData();
      for (const i of Object.keys(this.selectedFile)) {          
        formData.append("files", this.selectedFile[i]);
          }    
       axios.post("http://localhost:8000/route/file", formData)
        .then((res) => {
        this.filelocations = res.data.fileslocations;
        })
        .catch((err) => {
          console.log(err);
      });
    },

    convertstringtoarray(commaseparetdstring){
        var arrayvalue = [];
        let startindex = 0;
        var string =  "";
        for(let i=0; i<commaseparetdstring.length; i++){
            if(commaseparetdstring[i] == ","){
                for(let j=startindex; j<i; j++){
                     string+=commaseparetdstring[j];   
                }
                arrayvalue.push(string);
                startindex = i+1;
                string = "";
            }
        }
        return arrayvalue;
    },
     
    makeimagesofdata(commaseparetdstring, recipeid) {
       var datas = [];
       var arrayvalue =   this.convertstringtoarray(commaseparetdstring);
        for(let z=0; z<arrayvalue.length; z++){
           var obj = {};
           obj.image_link = arrayvalue[z];
           obj.recipes_id = recipeid;
           datas.push(obj);
       }
       return datas;
        
      },

    makeaingridentofdata(commaseparetdstring, recipeid) {
       var datas = [];
       var arrayvalue =   this.convertstringtoarray(commaseparetdstring);
        for(let z=0; z<arrayvalue.length; z++){
           var obj = {};
           obj.ingrident = arrayvalue[z];
           obj.recipes_id = recipeid;
           datas.push(obj);
       }
       return datas;
        
      },
      makeaingstepsofdata(commaseparetdstring, recipeid) {
        var datas = [];
        var arrayvalue =   this.convertstringtoarray(commaseparetdstring);
        for(let z=0; z<arrayvalue.length; z++){
           var obj = {};
           obj.step = arrayvalue[z];
           obj.recipes_id = recipeid;
           datas.push(obj);
        }
       return datas;
          },

  async  addfood() {
    await  this.onUploadFile();
      this.addrecipes({
        title: this.title,
        description: this.description,
        time: this.time,
        catagories: this.catagories,
        creatorsname: this.creatorsname,
        creatorsid: this.creatorsid,
      }).then((data) => {
        var objects = this.makeaingridentofdata(this.ingridents,data.data.insert_recipes_one.id);
        this.addingridents({objects:objects}).then((data)=>{  console.log(data) });
            objects = this. makeaingstepsofdata( this.steps,data.data.insert_recipes_one.id);
        this.addingsteps({objects:objects}).then((data)=>{ console.log(data) });
            objects = this.makeimagesofdata(this.filelocations,data.data.insert_recipes_one.id);
        this.addingimages({objects:objects}).then((data)=>{ console.log(data) });
          router.push("/adminpage");
      });
    },
  },
  setup() {
    const title = ref("en");
    const description = ref("en");
    const catagories = ref("en");
    const creatorsname = ref("en");
    const time = ref(0);
    const creatorsid = ref(0);
    const object = ref({});

    const { mutate: addrecipes } = useMutation(INSERT_RECIPES_ONE, () => ({
      variables: {
        title: title.value,
        description: description.value,
        time: time.value,
        catagories: catagories.value,
        creatorsname: creatorsname.value,
        creatorsid: creatorsid.value,
      },
    }));

    const { mutate: addingridents } = useMutation(INSERT_RECIPES_INGRIDENTS, () => ({
      variables: {
        objects: object.value
      },
    }));

    const { mutate: addingsteps } = useMutation(INSERT_RECIPES_STEPS, () => ({
      variables: {
        objects: object.value
      },
    }));
    const { mutate: addingimages } = useMutation(INSERT_RECIPES_IMAGES, () => ({
      variables: {
        objects: object.value
      },
    }));

    return {
      addrecipes,
      addingridents,
      addingsteps,
      addingimages
    };
  },
};
</script>
<template>
  <form
      action="#"
      method="post"
      enctype="multipart/form-data">

    <div class="grid md:grid-cols-2 md:gap-6 mx-12 mt-10">
      <div class="relative z-0 w-full mb-6 group">
        <input
          v-model="title"
          class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
          placeholder="title"
          required
        />
      </div>
      <div class="relative z-0 w-full mb-6 group">
        <input
          v-model="description"
          class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
          placeholder="discription"
          required
        />
      </div>
    </div>
    <div class="grid md:grid-cols-2 md:gap-6 mx-12">
      <div class="relative z-0 w-full mb-6 group">
        <input
          v-model="time"
          class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
          placeholder="time to prepare"
          required
        />
      </div>
      <div class="relative z-0 w-full mb-6 group">
        <input
          v-model="catagories"
          class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
          placeholder="food catagories"
          required
        />
      </div>
    </div>

    <div class="grid md:grid-cols-2 md:gap-6 mx-12">
      <div class="relative z-0 w-full mb-6 group">
        <textarea
          v-model="ingridents"
          rows="4"
          class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Ingridents comma separated values..."
        ></textarea>
      </div>
      <div class="relative z-0 w-full mb-6 group">
        <textarea
          v-model="steps"
          rows="4"
          class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Steps comma separated values..."
        ></textarea>
      </div>
    </div>

    <div class="grid md:grid-cols-2 md:gap-6 mx-12">
      <input
        class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
          id="file_input"
          type="file"
          name="files[]"
          ref="file"
          @change="onFileChange"
          multiple
      />
    </div>
    <button
      @click="addfood"
      type="button"
      class="text-white mt-2 mx-12 bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700"
      >
     Add Product
    </button>
  </form>
</template>
<style scoped></style>
