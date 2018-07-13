<template>
  <div class="Photohosting">
    <h1>{{ title }}</h1>

   <table align="center" border="1" width="100%" cellpadding="1" bgcolor="fadd73" bordercolor="black">
    <tr>
      <th>Имя</th>
      <th>Псевдоним</th>
      <th>Почта</th>
      <th>О себе</th>
      <th>День рождения</th>
      <th>Тема профиля</th>
    </tr>
    <tr v-for="contact in contact_list" v-bind:key="contact.name">
      <th>{{contact.name}}</th>
      <th>{{contact.nickname}}</th>
      <th>{{contact.email}}</th>
      <th>{{contact.description}}</th>
      <th>{{contact.birthday}}</th>
      <th>{{contact.photos}}</th>
      <th><button v-on:click="edit_contact(contact)" a title="Нажмите, чтобы выполнить редактирование">Редактировать пользователя</button> </th>
      <th><button v-on:click="remove_contact(contact)" a title="Нажмите, чтобы выполнить удаление">Удалить пользователя</button> </th>
    </tr>
   </table>

    <h3 v-if="edit_index == -1">Новый пользователь</h3>
   <form>
      <p>Имя: <input type="text" v-model="new_contact.name" a title="Введите имя" required></p>
      <p>Псевдоним <input type="text" v-model="new_contact.nickname" a title="Введите прозвище" required></p>
      <p>Почта <input type="email" v-model="new_contact.email" a title="Введите свою почту в формате ***@***.***" required pattern="^[a-zA-Z]+@+[a-zA-Z]+.+[a-zA-Z]"></p>
      <p>О себе <input type="text" v-model="new_contact.description" a title="Введите информацию о себе" required></p>
      <p>День рождения <input type="text" v-model="new_contact.birthday" a title="Введите свой день рождения" required></p>
      <p>Тема профиля <input type="text" v-model="new_contact.photos" a title="Добавьте тему профиля" required></p>
      <button v-if="edit_index == -1" v-on:click="add_new_contact" a title="Нажмите, чтобы выполнить добавление нового пользователя">Добавить пользователя</button>
      <button v-if="edit_index > -1" v-on:click="make_new_contact" a title="Нажмите, чтобы выполнить редактирование">Редактировать</button>
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
