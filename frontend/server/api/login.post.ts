// // ref: https://zenn.dev/howtelevision/articles/e8531bde04d7e4

// export default defineEventHandler(async (event) => {
//   const body = await readBody(event)
//   const target = body.target
//   const expiresIn = 60 * 60 * 24 * 5 * 1000

//   const options = {
//     maxAge: expiresIn,
//     httpOnly: true,
//     secure: true,
//   }

//   setCookie(event, "target", target, options)

//   return { message: "success" }
// })
