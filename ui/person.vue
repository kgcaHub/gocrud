<template>
  <div>
    <header>
      <h1 style="text-align: center">Person crud</h1>
    </header>
    <div class="columns">
      <div class="column is-half">
        <div class="field">
          <label class="label">Id</label>
          <div class="control">
            <input class="input" type="text">
          </div>
        </div>
        <div class="field">
          <label class="label">Name</label>
          <div class="control">
            <input class="input" type="text">
          </div>
        </div>
        <div class="field is-grouped">
          <p class="control">
            <a class="button">Create</a>
          </p>
          <p class="control">
            <a class="button">Update</a>
          </p>
          <p class="control">
            <a class="button">Delete</a>
          </p>
        </div>
      </div>
      <div class="column is-half">
        <table>
          <thead>
            <tr>
              <th>Id</th>
              <th>Name</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>123</td>
              <td>Klinsmann</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
<script>
export default {
    name:"Person",
    props:{},
    data(){
        return {};
    },
    created(){

    },
    methods:{

    }
};
</script>

