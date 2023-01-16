import fs from 'fs'
import path from 'path'

const createRobotsTxt = () => {
  if (process.env.ROBOTS_TXT_ALLOW === 'true') {
    return `User-agent: *\nAllow: /`
  } else {
    return `User-agent: *\nDisallow: /`
  }
}

export const generateRobotsTxt = () => {
  fs.writeFileSync(path.join(__dirname, '..', 'public', 'robots.txt'), createRobotsTxt())
}
