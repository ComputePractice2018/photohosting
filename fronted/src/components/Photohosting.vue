<template>
  <div class="Photohosting">
    <h1>{{ title }}</h1>

   <table>
    <tr>
      <th>Имя</th>
      <th>Псевдоним</th>
      <th>Почта</th>
      <th>О себе</th>
      <th>День рождения</th>
      <th>Фотографии</th>
    </tr>
    <tr v-for="contact in contact_list" v-bind:key="contact.name">
      <td>{{contact.name}}</td>
      <td>{{contact.nickname}}</td>
      <td>{{contact.email}}</td>
      <td>{{contact.description}}</td>
      <td>{{contact.birthday}}</td>
      <td>{{contact.photos}}</td>
      <td><button v-on:click="edit_contact(contact)">Редактировать пользователя</button> </td>
      <td><button v-on:click="remove_contact(contact)">Удалить пользователя</button> </td>
    </tr>
   </table>

    <h3 v-if="edit_index == -1">Новый пользователь</h3>
   <form>
      <p>Имя: <input type="text" v-model="new_contact.name"></p>
      <p>Псевдоним <input type="text" v-model="new_contact.nickname"></p>
      <p>Почта <input type="text" v-model="new_contact.email"></p>
      <p>О себе <input type="text" v-model="new_contact.description"></p>
      <p>День рождения <input type="text" v-model="new_contact.birthday"></p>
      <p>Фотографии <input type="text" v-model="new_contact.photos"></p>
      <button v-if="edit_index == -1" v-on:click="add_new_contact">Добавить пользователя</button>
      <button v-if="edit_index > -1" v-on:click="make_new_contact">Редактировать</button>
   </form>
</div>
</template>

<script>
export default {
  name: 'Photohosting',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      contact_list: [
        {
          'name': 'Марат',
          'nickname': 'Saffin',
          'email': 'Почта',
          'description': 'О себе',
          'birthday': 'День рождения',
          'photos': 'Фотографии:'
        },
        {
          'name': 'Полина',
          'nickname': 'ladowska',
          'email': 'Почта',
          'description': 'О себе',
          'birthday': 'День рождения',
          'photos': 'Фотографии:'
        }
      ],
      new_contact:
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
  methods: {
    add_new_contact: function () {
      this.contact_list.push(this.new_contact)
    },
    remove_contact: function (item) {
      this.contact_list.splice(this.contact_list.indexOf(item), 1)
    },
    edit_contact: function (item) {
      this.edit_index = this.contact_list.indexOf(item)
      this.new_contact = this.contact_list[this.edit_index]
    },
    make_new_contact: function () {
      this.edit_index = -1
      this.new_contact = {
        'name': '',
        'nickname': '',
        'email': '',
        'description': '',
        'birthday': '',
        'photos': ''
      }
    }
  }
}
</script>

<style scoped>
</style>
