import org.graalvm.nativeimage.IsolateThread;
import org.graalvm.nativeimage.c.function.CEntryPoint;
import org.graalvm.nativeimage.c.function.CFunctionPointer;
import org.graalvm.nativeimage.c.function.InvokeCFunctionPointer;
import org.graalvm.nativeimage.c.type.CCharPointer;
import com.oracle.svm.core.c.CConst;
import org.graalvm.nativeimage.c.type.CTypeConversion;

import java.time.Instant;

public final class LibJavaCgo {
    interface AllocatorFn extends CFunctionPointer
    {
        @InvokeCFunctionPointer
        CCharPointer call(long size);
    }

    @CEntryPoint(name = "java_cgo_str")
    public static @CConst CCharPointer javaCgoStr(IsolateThread thread, AllocatorFn alloc, CCharPointer str) {
        System.out.println(CTypeConversion.toJavaString(str));
        byte[] b = Instant.now().toString().getBytes();
        CCharPointer a =  alloc.call(b.length + 1);
        for (int i = 0; i < b.length; i++) {
            a.write(i, b[i]);
        }
        a.write(b.length, (byte) 0);
        return a;
    }
}
