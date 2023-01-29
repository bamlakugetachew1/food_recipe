<script>
import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import { LOGIN_ACTIONS_MUTATION } from "../graphql/graphql";
import router from "../router/index";
import { Field, Form } from "vee-validate";

export default {
  name: "loginpage",
  components: {
    Field,
    Form,
  },
  data() {
    return {
      name: "",
      password: "",
      message: "",
    };
  },
  methods: {
    login() {
      this.loginuser({ name: this.name, password: this.password }).then(
        (data) => {
          this.message = data.data.login.success;
          if (data.data.login.success == "true") {
            localStorage.setItem("token", data.data.login.accessToken);
            localStorage.setItem("userid",data.data.login.id);
            localStorage.setItem("username",data.data.login.name);
            router.push({
              name: "adminpage",
              params: { userid: data.data.login.id },
            });
          }
        }
      );
    },
    isRequired(value) {
      return value ? true : "field is required";
    },
  },
  setup() {
    const name = ref("");
    const password = ref("");
    const { mutate: loginuser } = useMutation(LOGIN_ACTIONS_MUTATION, () => ({
      variables: {
        name: name.value,
        password: password.value,
      },
    }));
    return {
      loginuser,
    };
  },
};
</script>

<template>
  <section class="bg-gray-50 dark:bg-gray-900">
    <div
      class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
    >
      <div
        class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700"
      >
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <p class="mb-2 text-center font-serif capitalize">
            {{ this.message }}
          </p>
          <h1
            class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white"
          >
            Sign in to your account
          </h1>

          <Form v-slot="{ errors }" class="space-y-4 md:space-y-6">
            <label
              for="name"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Your name</label
            >
            <Field
              name="field1"
              :rules="isRequired"
              v-model="name"
              placeholder="bamlaku"
              class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
            <span>{{ errors.field }}</span>
            <label
              for="password"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Password</label
            >
            <Field
              name="field2"
              type="password"
              :rules="isRequired"
              v-model="password"
              placeholder="........"
              class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
            <span>{{ errors.field }}</span>
            <button
              @click="login"
              type="button"
              class="text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700"
            >
              Login
            </button>
          </Form>
        </div>
      </div>
    </div>
  </section>
</template>
<style scoped></style>
