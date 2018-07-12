<template>
  <div class="Photohosting">
     <b-jumbotron bg-variant="info" text-variant="white" border-variant="dark">
      <template slot="header">
        {{ title}}
      </template>
    </b-jumbotron>

    <h3 v-if="error">Ошибка: {{error}}</h3>

    <table class="table">
      <thead>
      <tr>
        <th>name</th>
        <th>nickname</th>
        <th>email</th>
        <th>description</th>
        <th>birthday</th>
        <th>photos</th>
        <th></th>
        <th></th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="profile in profile" v-bind:key="profile.id">
        <td>{{profile.name}}</td>
        <td>{{profile.nickname}}</td>
        <td>{{profile.email}}</td>
        <td>{{profile.description}}</td>
        <td>{{profile.birthday}}</td>
        <td>{{profile.photos}}</td>
        <td><b-button variant="warning" v-on:click="edit_profile(profile)">Редактировать профиль</b-button></td>
        <td><b-button variant="danger" v-on:click="remove_profile(profile)">Удалить профиль</b-button></td>
      </tr>
      </tbody>
    </table>

    <b-card>
    <b-form>
      <h3 v-if="edit_index == -1">Новый профиль</h3>
      <b-form-group label="name:">
        <b-form-input v-model="new_profile.name" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="nickname:">
        <b-form-input v-model="new_profile.surname" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="email:">
        <b-form-input v-model="new_profile.phone" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="description:">
        <b-form-input v-model="new_profile.email" type="email" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="birthday:">
        <b-form-input v-model="new_profile.github" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="photos:">
        <b-form-input v-model="new_profile.github" required>
        </b-form-input>
      </b-form-group>
      <b-button variant="success" v-if="edit_index == -1" v-on:click="add_new_profile">Добавить профиль</b-button>
      <b-button variant="success" v-if="edit_index > -1" v-on:click="end_of_profile">Закончить редактирование</b-button>
    </b-form>
    </b-card>
  </div>
</template>

<script>
const axios = require('axios')
export default {
  name: 'Photohosting',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      error: '',
      profile_list: [],
      new_profile:
        {
          'id': 0,
          'name': '',
          'nickname': '',
          'email': '',
          'description': '',
          'birthday': '',
          'photos': ''
        }
    }
  },
  mounted: function () {
    this.get_profiles()
  },
  methods: {
    get_profiles: function () {
      this.error = ''
      const url = '/api/Photohosting/profiles'
      axios.get(url).then(response => {
        this.profile_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_profile: function () {
      this.error = ''
      const url = '/api/Photohosting/profiles'
      axios.post(url, this.new_profile).then(response => {
        this.profile_list.push(this.new_profile)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_profile: function (item) {
      this.error = ''
      const url = '/api/Photohosting/profile/' + item.id
      axios.delete(url).then(response => {
        this.profile_list.splice(this.profile_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_profile: function (item) {
      this.edit_index = this.profile_list.indexOf(item)
      this.new_profile = this.profile_list[this.edit_index]
    },
    end_of_edition: function () {
      this.error = ''
      const url = '/api/Photohosting/profile/' + this.new_profile.id
      axios.put(url, this.new_profile).then(response => {
        this.edit_index = -1
        this.new_profile = {
          'id': 0,
          'name': '',
          'nickname': '',
          'email': '',
          'description': '',
          'birthday': '',
          'photos': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>

<style scoped>

</style>
