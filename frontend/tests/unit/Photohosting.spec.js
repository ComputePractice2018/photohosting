import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import Photohosting from '@/components/Photohosting.vue'

describe('Photohosting.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Name'
    const wrapper = shallowMount(Photohosting, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(Photohosting, {})
    expect(wrapper.text()).to.include('name')
    expect(wrapper.text()).to.include('nickname')
    expect(wrapper.text()).to.include('email')
    expect(wrapper.text()).to.include('description')
    expect(wrapper.text()).to.include('birthday')
    expect(wrapper.text()).to.include('photos')
  })
})
