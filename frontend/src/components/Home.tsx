import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบแจ้งซ่อม</h1>
        <h4>Requirements</h4>
        <p>
        ระบบแจ้งซ่อมของห้องพักเป็นระบบที่ให้ผู้ใช้ซึ่งเป็นผู้เช่าสามารถ login เข้ามาเพื่อที่จะแจ้งซ่อมอุปกรณ์หรือเฟอร์นิเจอร์ที่มีการชำรุดเสียหาย 
        ผู้เช่าจะสามารถแจ้งซ่อมได้มากกว่า 1 รายการ ในการแจ้งซ่อมนั้นสามารถแจ้งซ่อมได้เฉพาะอุปกรณ์ภายในห้องพักตามในใบสัญญาเช่าที่ผู้เช่าได้เช่าไว้ในระบบเท่านั้น 
        โดยการแจ้งซ่อมสามารถระบุอุปกรณ์ ประเภท หรือแจ้งหมายเหตุเพิ่มเติมเข้าไปได้ ผู้เช่าสามารถเพิ่มได้ด้วยตัวเอง เมื่อทำการเพิ่มรายการการแจ้งซ่อมเรียบร้อยแล้ว 
        ระบบก็จะบันทึกและแสดงรายการแจ้งซ่อมขึ้นที่หน้าจอ
        </p>
      </Container>
    </div>
  );
}
export default Home;
