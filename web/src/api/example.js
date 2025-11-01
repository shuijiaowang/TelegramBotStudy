import service from '@/utils/request.js'

export const text = (data) => {
    return service({
        url: '/example/test',
        method: 'post',
        data: data
    })
}
