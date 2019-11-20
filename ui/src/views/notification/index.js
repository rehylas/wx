import {   notification } from 'antd';


export function openNotification_succ(msg, desp)  {    

    notification.success({message:  msg, description:desp  });

};

export function openNotification_err(msg, desp)  {    

    notification.error({message:  msg, description:desp  });

};

export function openNotification_warn(msg, desp)  {    

    notification.warn({message:  msg, description:desp  });

};

export function openNotification_info(msg, desp)  {    

    notification.info({message:  msg, description:desp  });

};



// export default openNotification;

 

 