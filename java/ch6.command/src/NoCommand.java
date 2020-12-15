public class NoCommand implements Command {

    @Override
    public void execute() {
        System.out.println("아무 기능이 없음");
    }

}
