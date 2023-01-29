import router from "../router/index"
export function checkuserloggedin() {
    if(localStorage.getItem("token")==null){
         router.push("/loginpage")    
      }
  }

  