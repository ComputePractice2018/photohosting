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
  it('renders Имя', () => {
    const wrapper = shallowMount(Photohosting, {})
    expect(wrapper.text()).to.include('Name')
    expect(wrapper.text()).to.include('Nickname')
    expect(wrapper.text()).to.include('Email')
    expect(wrapper.text()).to.include('description')
    expect(wrapper.text()).to.include('birthday')
    expect(wrapper.text()).to.include('photos')
})
