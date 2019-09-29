# 写一个转发器

主要有两个组件：
- 判断request url, 如果url是 get_th ，才会 modifyResponse?? 
- 写modifyResponse.
- 怎样改写 res.body? type io.ReadCloser

从 http.Response 里找写 response的方法。