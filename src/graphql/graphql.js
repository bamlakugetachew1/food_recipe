import gql from "graphql-tag";
export var SIGNUP_ACTIONS_MUTATION = gql`
mutation  signup($name:String!,$password:String!){
     signup(name:$name,password:$password){
        accessToken
        success
     }
     }
`
export var LOGIN_ACTIONS_MUTATION = gql`
mutation  login($name:String!,$password:String!){
     login(name:$name,password:$password){
        accessToken
        success
        id
        name
  }
}
`

export const RETRIVE_ALL_RECIPES_QUERY = gql`
query myquery ($limit:Int!){
  recipes(limit: $limit) {
    id
    title
    images {
      image_link
      fullimagelink
    }
  }
}
`

export const RETRIVE_AGGREGATE_RATING_RECIPES_QUERY = gql`
query($recipeid:Int!){
  count_ratings(args: {recipeid: $recipeid}) {
    id
    rate
  }
  count_ratings_aggregate(args: {recipeid: $recipeid}) {
    aggregate {
      count
      avg {
        rate
      }
      min {
        rate
      }
      max {
        rate
      }
    }
  }
}
`


export const FILTER_BY_TIME = gql`
query myquery($limit: Int!) {
  recipes(limit: $limit, order_by: {time: asc}) {
    id
    title
    images {
      image_link
      fullimagelink
    }
  }
}
`

export const FILTER_BY_TITLE = gql`
query myquery($limit: Int!) {
  recipes(limit: $limit, order_by: {title: asc}) {
    id
    title
    images {
      image_link
      fullimagelink
    }
  }
}
`

export const BROWSEBY_TITLE_CREATORSNAME_OR_CATAGORY = gql`
  query myquery($limit: Int!, $search: String!) {
  recipes(limit: $limit, where: {   _or: [
        {title: {_like: $search}},
        {creatorsname: {_like: $search}},
        {catagories: {_like: $search},

        }
      ]}) {
    id
    title
    images {
      image_link
      fullimagelink
    }
  }
}
`


export const RETRIVE_COMMENTS = gql`
  query comments($recipes_id:Int!) {
    comment(where: {recipes_id: {_eq: $recipes_id}}) {
      comment
    }
}
`
export const INSERT_RATINGS = gql`
mutation($rate:Int!,$user:Int!,$recipeid:Int!){
  insert_rating_one(object:{rate:$rate,users_id:$user,recipes_id:$recipeid}){
    rate
  }
}
`




 export const CREATE_COMMENTS = gql`
   mutation($comment:String!,$user:Int!,$recipeid:Int!){
  insert_comment_one(object:{comment:$comment,users_id:$user,recipes_id:$recipeid}){
    comment
 }
}
`




export const RETRIVE_SINGLE_RECIPES_QUERY = gql`
query myquery ($id:Int!) {
  recipes_by_pk(id: $id){
    id
    title
    description
    catagories
    time
    creatorsname
    creatorsid
    ingridents{
      id
      ingrident
    }
    steps{
      step
      id
    }
    images{
      image_link
      id
    }
  }
}

`
export var DELETE_RECIPES = gql`
  mutation ($id:Int!) {
    delete_recipes_by_pk(id: $id){
    id
  }
}
`

export var CHECK_ALERADY_LIKED = gql`
query likes ($users_id:Int!,$recipes_id:Int!) {
  like(where: {users_id: {_eq:$users_id}, recipes_id: {_eq: $recipes_id}}){
     id
  }
}
`
export var ADD_USER_LIKES = gql`
mutation ($userid:Int!,$recipeid:Int!) {
  insert_like(objects:[{users_id:$userid,recipes_id:$recipeid}]){
    affected_rows
    returning{
      users_id
    }
  }
}
`




export var INSERT_RECIPES_ONE = gql`
mutation($title:String!,$catagories:String!,$description:String!,$time:Int!,$creatorsname:String!,$creatorsid:Int!){
   insert_recipes_one(object:{title:$title,catagories:$catagories,description:$description,time:$time,creatorsname:$creatorsname,creatorsid:$creatorsid}){
     creatorsname
     creatorsid
     id
  }
}
`
export var INSERT_RECIPES_IMAGES = gql`
mutation($objects:[images_insert_input!]!){
  insert_images(objects:$objects){
    affected_rows
  }
}
`
export var INSERT_RECIPES_INGRIDENTS = gql`
mutation($objects:[ingridents_insert_input!]!){
  insert_ingridents(objects:$objects){
    affected_rows
  }
}
`

export var INSERT_RECIPES_STEPS = gql`
mutation($objects:[steps_insert_input!]!){
  insert_steps(objects:$objects){
    affected_rows
  }
}
`




