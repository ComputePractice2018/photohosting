import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import Photohosting from '@/components/Photohosting.vue'

describe('Photohosting.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(Photohosting, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(Photohosting, {})
    expect(wrapper.text()).to.include('Имя')
    expect(wrapper.text()).to.include('Прозвище')
    expect(wrapper.text()).to.include('Почта')
    expect(wrapper.text()).to.include('О себе')
    expect(wrapper.text()).to.include('День Рождения')
    expect(wrapper.text()).to.include('Фотографии')
  })
})
