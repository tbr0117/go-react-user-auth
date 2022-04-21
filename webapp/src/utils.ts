export const CreateCookie = (cookieName:string, cookieValue:string, hourToExpire:number) => {
    const date:Date = new Date()
    date.setTime(date.getTime() + hourToExpire * 60 * 60 * 1000,)
    document.cookie = `${cookieName} = ${cookieValue}; expires = ${date.toUTCString()}`
  }
  
  export const DeleteCookie = (name:string) =>
    (document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;')
  