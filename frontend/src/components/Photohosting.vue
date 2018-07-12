<template>
  <div class="Photohosting">
    <h1>{{ title }}</h1>

    <h3 v-if="error">Ошибка: {{error}}</h3>

    <table>
      <tr>
        <th>name</th>
        <th>nickname</th>
        <th>email</th>
        <th>description</th>
        <th>birthday</th>
        <th>photos</th>
      </tr>
      <tr v-for="profile in profile_list" v-bind:key="profile.email">
        <td>{{profile.name}}</td>
        <td>{{profile.nickname}}</td>
        <td>{{profile.email}}</td>
        <td>{{profile.description}}</td>
        <td>{{profile.birthday}}</td>
        <td>{{profile.photos}}</td>
        <td><button v-on:click="edit_profile(profile)">Редактировать контакт</button></td>
        <td><button v-on:click="remove_profile(profile)">Удалить контакт</button></td>
      </tr>
    </table>

    <h3 v-if="edit_index == -1">Новый контакт</h3>
    <form>
      <p>Имя: <input type="text" v-model="new_profile.name"></p>
      <p>Фамилия: <input type="text" v-model="new_profile.nickname"></p>
      <p>Телефон: <input type="text" v-model="new_profile.email"></p>
      <p>description: <input type="text" v-model="new_profile.description"></p>
      <p>Gihub: <input type="text" v-model="new_profile.birthday"></p>
      <p>Gihub: <input type="text" v-model="new_profile.photos"></p>
      <button v-if="edit_index == -1" v-on:click="add_new_profile">Добавить контакт</button>
      <button v-if="edit_index > -1" v-on:click="end_of_edition">Закончить редактирование</button>
    </form>
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
      const url = '/api/Photohosting/profile/' + this.profile_list.indexOf(item)
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
      const url = '/api/Photohosting/profile/' + this.edit_index
      axios.put(url, this.new_profile).then(response => {
        this.edit_index = -1
        this.new_profile = {
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
